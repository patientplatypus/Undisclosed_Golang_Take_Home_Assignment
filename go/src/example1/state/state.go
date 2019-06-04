package state

import()

type Pet struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Tag       string   `json:"tag"`
}

var Petarray = []Pet{
	{ID: 1, Name: "Dog", Tag: "available"},
	{ID: 2, Name: "Cat", Tag: "pending"},
}
