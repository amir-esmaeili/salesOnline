package forms

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
