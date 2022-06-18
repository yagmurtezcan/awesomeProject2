package entity

type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
