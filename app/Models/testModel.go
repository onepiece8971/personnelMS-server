package Models

type Users struct {
	Id int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func (u Users) GetUser() (user Users) {
    db.Where("id = ?", u.Id).First(&user)
    return
}