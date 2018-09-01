package webserver

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/bah2830/brentahughes.com/repo"
)

var templatePath = "templates/"

type Webserver struct {
	repoClient *repo.RepoClient
}

type Page struct {
	Title                string
	Repos                []Repo
	Contributions        []Repo
	Name                 string
	Email                string
	PhoneNumber          string
	SocialIcons          []SocialIcon
	ProjectSource        string
	ProjectLocation      string
	ProjectLocationLower string
}

type Repo struct {
	Name string
	URL  string
	Repo string
}
type SocialIcon struct {
	Site string
	URL  string
}

func GetWebserver(c *repo.RepoClient) *Webserver {
	return &Webserver{
		repoClient: c,
	}
}

func (w *Webserver) Start() {
	http.HandleFunc("/", w.indexHandler)

	// Setup file server for html resources
	fs := http.FileServer(http.Dir("content"))
	http.Handle("/content/", http.StripPrefix("/content/", fs))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setUserDetails(p *Page) {
	p.Title = viper.GetString("site_title")
	p.Name = viper.GetString("name")
	p.Email = viper.GetString("email")
	p.PhoneNumber = viper.GetString("phone")
	p.ProjectSource = viper.GetString("project_source")

	urlParts, _ := url.Parse(p.ProjectSource)
	projectIcon := strings.TrimPrefix(urlParts.Host, "www.")
	parts := strings.Split(projectIcon, ".")
	projectIcon = parts[0]
	p.ProjectLocation = strings.Title(projectIcon)
	p.ProjectLocationLower = strings.ToLower(projectIcon)
}

func getSocialIcons() []SocialIcon {
	icons := []SocialIcon{}
	links := viper.GetStringSlice("social_links")
	for _, link := range links {
		parts, err := url.Parse(link)
		if err != nil {
			fmt.Println(err)
		}

		host := strings.Replace(parts.Host, "www.", "", -1)
		hostParts := strings.Split(host, ".")

		icon := SocialIcon{
			Site: hostParts[0],
			URL:  link,
		}

		icons = append(icons, icon)
	}

	return icons
}

func (s *Webserver) indexHandler(w http.ResponseWriter, r *http.Request) {
	originalRepos := make([]Repo, 0)
	contributions := make([]Repo, 0)
	repos, _ := s.repoClient.GetRepos(false)
	for _, r := range repos {
		repo := Repo{
			URL:  r.GetURL(),
			Name: r.GetName(),
			Repo: r.GetRepo(),
		}

		if r.IsContribution() {
			contributions = append(contributions, repo)
		} else {
			originalRepos = append(originalRepos, repo)
		}
	}

	p := Page{
		Repos:         originalRepos,
		Contributions: contributions,
	}

	p.SocialIcons = getSocialIcons()

	setUserDetails(&p)

	templates := template.Must(template.ParseFiles(templatePath + "index.html"))
	templates.ExecuteTemplate(w, "index.html", p)
}
