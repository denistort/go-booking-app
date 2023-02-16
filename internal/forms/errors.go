package forms

type errors map[string][]string

// Add adds an error message for a given field
func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

// Get return the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
