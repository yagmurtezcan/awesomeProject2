package dto

type UserDTO struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (UserDTO) TableName() string {
	return "user"
}
