package people

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) create(p *PeopleEntity) error {
	var strStack sql.NullString
	if p.Stack != nil {
		strStack.String = strings.Join(p.Stack, ",")
		strStack.Valid = true
	}
	_, err := r.db.Exec("INSERT INTO people (id, name, nickname, birthdate, stack, search) VALUES ($1, $2, $3, $4, $5, $6)", p.ID, p.Name, p.Nickname, p.BirthDate, strStack, p.Search)
	if err != nil {
		return ErrNicknameTaken
	}

	return nil
}

func (r *repository) readByID(id string) (*PeopleEntity, error) {
	row := r.db.QueryRow("SELECT id, name, nickname, birthdate, stack, search FROM people WHERE id = $1", id)
	if err := row.Err(); err != nil {
		fmt.Println(err)

		if sql.ErrNoRows == err {
			return nil, ErrNoPeopleFound
		}

		return nil, ErrInternal
	}

	var p PeopleEntity
	var strStack sql.NullString
	err := row.Scan(&p.ID, &p.Name, &p.Nickname, &p.BirthDate, &strStack, &p.Search)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, ErrNoPeopleFound
		}

		return nil, ErrInternal
	}

	if strStack.Valid {
		p.Stack = strings.Split(strStack.String, ",")
	}

	return &p, nil
}

func (r *repository) searchByTerm(term string) ([]*PeopleEntity, error) {
	rows, err := r.db.Query("SELECT id, name, nickname, birthdate, stack, search FROM people WHERE search LIKE $1", "%"+term+"%")
	if err != nil {
		return nil, err
	}

	var peoples []*PeopleEntity
	for rows.Next() {
		var p PeopleEntity
		var strStack sql.NullString
		err := rows.Scan(&p.ID, &p.Name, &p.Nickname, &p.BirthDate, &strStack, &p.Search)
		if err != nil {
			return nil, ErrInternal
		}

		if strStack.Valid {
			p.Stack = strings.Split(strStack.String, ",")
		}

		peoples = append(peoples, &p)
	}

	return peoples, nil
}

func (r *repository) count() (int64, error) {
	row := r.db.QueryRow("SELECT COUNT(*) FROM people")
	if err := row.Err(); err != nil {
		return 0, err
	}

	var count int64
	err := row.Scan(&count)
	if err != nil {
		return 0, ErrInternal
	}

	return count, nil
}
