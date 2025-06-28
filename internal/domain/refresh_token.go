package domain

type RefreshToken string

func (r RefreshToken) String() string {
	return string(r)
}

func (r RefreshToken) IsValid() bool {
	if len(r) == 0 {
		return false
	}

	return true
}
