package domain

type Image struct {
	ID       string
	RepoTags []string
	Size     int64
}
