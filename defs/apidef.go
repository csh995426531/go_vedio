package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VedioInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}
