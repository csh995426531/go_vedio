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
