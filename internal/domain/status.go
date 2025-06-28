package domain

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

var statuses = map[Status]struct{}{
	Active:   {},
	Inactive: {},
}

type Status string

func (s Status) IsValid() bool {
	if len(s) == 0 {
		return false
	}

	if _, ok := statuses[s]; !ok {
		return false
	}

	return true
}
