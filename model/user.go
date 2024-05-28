package model

type User struct {
	Id          uint   `json:"id"`
	VoorNaam    string `json:"voornaam"`
	FamilieNaam string `json:"familienaam"`
}

type PostUser struct {
	VoorNaam    string `json:"voornaam"`
	FamilieNaam string `json:"familienaam"`
}

type UserUri struct {
	ID uint `uri:"id" binding:"required,number"`
}

type UpdateUser struct {
	VoorNaam    string `json:"voornaam"`
	FamilieNaam string `json:"familienaam"`
}
