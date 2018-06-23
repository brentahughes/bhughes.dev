package projects

// Client connects to the remote resource
type Client interface {
	GetPrivateRepos() []Repo
	GetPublicRepos() []Repo
}

type Repo interface {
	GetName() string
	GetURL() string
	GetSource() string
}
