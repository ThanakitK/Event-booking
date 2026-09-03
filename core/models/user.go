package models

type RepoCreateUserModel struct {
	Name         string `bson:"name"`
	User         string `bson:"user"`
	PasswordHash string `bson:"password_hash"`
}

type SrvCreateUserModel struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}
