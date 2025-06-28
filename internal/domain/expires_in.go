package domain

type ExpiresIn int64

func (e ExpiresIn) IsValid() bool {
	if e == 0 {
		return false
	}

	return true
}
