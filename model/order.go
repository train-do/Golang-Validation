package model

type Order struct {
	Id                    int
	DestinationScheduleId int    `validate:"required"`
	Name                  string `validate:"required"`
	Email                 string `validate:"required,email"`
	EmailConfirm          string `validate:"required,email"`
	Phone                 string `validate:"required"`
	Message               string
	Rating                float32 `validate:"required"`
}
