package forms

import (
	"mime/multipart"
)

type SignUpForm struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LogInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileForm struct {
	Address string `form:"address"`
	Phone   string `form:"phone"`
	ProfileImage *multipart.FileHeader `form:"profile_image"`
}
