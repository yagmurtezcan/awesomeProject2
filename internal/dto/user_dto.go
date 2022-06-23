package dto

type UserDTO struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" validate:"required,gte=2,lte=50"`
	LastName  string `json:"last_name" validate:"required,gte=2,lte=50"`
	Email     string `json:"email" validate:"required,email,lte=70"`
}

func (UserDTO) TableName() string {
	return "user"
}
