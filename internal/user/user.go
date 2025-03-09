package user

type User struct {
	Status Status
}

type Status string

func (u *User) GetStatus() Status {
	return u.Status
}

func (u *User) SetStatus(status Status) {
	u.Status = status
}
