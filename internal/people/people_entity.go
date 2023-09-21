package people

import "time"

type PeopleEntity struct {
	ID        string
	Name      string
	Nickname  string
	BirthDate time.Time
	Stack     []string
	Search    string
}
