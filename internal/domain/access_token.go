package domain

type AccessToken string

func (a AccessToken) String() string {
	return string(a)
}

func (a AccessToken) IsValid() bool {
	if len(a) == 0 {
		return false
	}

	return true
}
