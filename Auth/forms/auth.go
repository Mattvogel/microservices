package forms

//Token ...
type Token struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}

//LoginFormWrapper ...
type LoginFormWrapper struct {
	User LoginForm `json:"user"`
}

//LoginForm ...
type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

//RegisterFormWrapper ...
type RegisterFormWrapper struct {
	User RegisterForm `json:"user"`
}

//RegisterForm ...
type RegisterForm struct {
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
