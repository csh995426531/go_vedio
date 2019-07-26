package defs

//UserCredential 用户
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//VedioInfo 视频信息
type VedioInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

//Comment 评论
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

//SimpleSession session
type SimpleSession struct {
	Username string //login name
	TTL      int64
}
