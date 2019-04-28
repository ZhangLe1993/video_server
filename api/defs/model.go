package defs

//request
type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

//Data Model

type Video struct {
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

type SessionVO struct {
	UserName string

	TTL int64
}
