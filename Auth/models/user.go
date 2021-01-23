package models

import (
	"errors"

	"Auth/db"
	"Auth/forms"

	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

//UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

//Login ...
func (m UserModel) Login(form forms.LoginFormWrapper) (user User, token Token, err error) {

	err = db.GetDB().SelectOne(&user,
		"SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1",
		form.User.Email)

	if err != nil {
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.User.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, errors.New("Invalid password")
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

//Register ...
func (m UserModel) Register(form forms.RegisterFormWrapper) (user User, token Token, err error) {
	getDb := db.GetDB()

	//Check if the user exists in database
	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.User.Email)

	if err != nil {
		return user, token, err
	}

	if checkUser > 0 {
		return user, token, errors.New("User already exists")
	}

	bytePassword := []byte(form.User.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err) //Something really went wrong here...
	}

	//Create the user and return back the user ID
	err = getDb.QueryRow("INSERT INTO public.user(email, password, name) VALUES($1, $2, $3) RETURNING id",
		form.User.Email,
		string(hashedPassword), form.User.Name).Scan(&user.ID)

	user.Name = form.User.Name
	user.Email = form.User.Email
	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, err
}

//One ...
func (m UserModel) One(userID int64) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT email, name FROM public.user WHERE id=$1", userID)
	return user, err
}

//Profile ...
type Profile struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//GetProfile ...
func (m UserModel) GetProfile(userID int64) (user Profile, err error) {

	prof := Profile{}

	err = db.GetDB().SelectOne(&prof, "Select id, email, name FROM public.user WHERE id=$1", userID)
	user = prof

	return user, err
}
