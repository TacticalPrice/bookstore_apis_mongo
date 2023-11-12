package models

type Book struct {
	ID string `json: "id, omitempty" bson: "_id, omitempty"`
	Title string `json:"title,omitempty" bson :"title, omitempty"`
	Author string `json: "author,omitempty" bson: "author, omitempty"`
}