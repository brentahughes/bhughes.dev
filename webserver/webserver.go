package webserver

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/bah2830/personal-website/github"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Page struct {
	Title         string
	Repos         []github.Repo
	Contributions []github.Repo
	Name          string
	Email         string
	PhoneNumber   string
	Address       string
	SocialIcons   []SocialIcon
}

type SocialIcon struct {
	Site string
	URL  string
}

func Start() {
	fmt.Println("Starting webserver...")

	http.HandleFunc("/", indexHandler)

	// Setup file server for html resources
	fs := http.FileServer(http.Dir("content"))
	http.Handle("/content/", http.StripPrefix("/content/", fs))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setUserDetails(p *Page) {
	p.Name = viper.GetString("name")
	p.Email = viper.GetString("email")
	p.PhoneNumber = viper.GetString("phone_number")
	p.Address = viper.GetString("address")

}

func getSocialIcons() []SocialIcon {
	icons := []SocialIcon{}

	links := viper.GetStringSlice("social_links")
	for _, link := range links {
		parts := strings.Split(link, "|")

		icon := SocialIcon{
			Site: parts[0],
			URL:  parts[1],
		}

		icons = append(icons, icon)
	}

	return icons
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Repos:         github.GetRepos(),
		Contributions: github.GetContributions(),
	}

	p.SocialIcons = getSocialIcons()

	setUserDetails(&p)

	templates.ExecuteTemplate(w, "index.html", p)
}
