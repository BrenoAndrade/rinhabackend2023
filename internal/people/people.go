package people

type People struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"nome"`
	Nickname  string   `json:"apelido"`
	BirthDate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

func (p *People) Validate() error {
	if p.Name == "" || len(p.Name) > 100 {
		return ErrInvalidPeople
	}

	for i, r := range p.BirthDate {
		if i == 4 || i == 7 {
			if r != '-' {
				return ErrInvalidPeople
			}

			continue
		}

		if r < '0' || r > '9' {
			return ErrInvalidPeople
		}
	}

	if p.Nickname == "" || len(p.Nickname) > 32 {
		return ErrInvalidPeople
	}

	if p.Stack != nil {
		for _, s := range p.Stack {
			if s == "" || len(s) > 32 {
				return ErrInvalidPeople
			}
		}
	}

	return nil
}
