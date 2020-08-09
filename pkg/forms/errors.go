package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(filed string) string {
	es := e[filed]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
