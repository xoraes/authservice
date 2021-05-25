package middleware

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"errors"
	_ "github.com/lib/pq" // postgres golang driver
	"github.com/sirupsen/logrus"
	"github.com/xoraes/dappauth/database"
	"github.com/xoraes/dappauth/models"
	"net/http" // used to access the request and response object of the api
)

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// create an empty user of type models.User
	var user models.User

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		logrus.Errorf("Unable to decode the request body.  %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if !user.Valid() {
		http.Error(w, "Invalid User", http.StatusBadRequest)
		return
	}
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//set the hash
	hashedPass, err := user.HashPassword()
	if err != nil {
		logrus.Errorf("Unable to create hash of the given password.  %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// call insert user function and pass the user
	err = database.InsertUser(user.Email, hashedPass, user.FirstName, user.LastName)
	if err != nil {
		logrus.Errorf("Unable to insert user.  %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	tkn, err := models.CreateToken(user.Email)
	if err != nil {
		logrus.Errorf("Unable to create token.  %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	// format a response object
	res := models.TokenResponse{
		Token: tkn,
	}
	// send the response
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	email, err := getEmailFromTokenHeader(r)
	if err != nil || email == "" {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	users, err := database.GetAllUsers()
	if err != nil {
		logrus.Errorf("Unable to get all user. %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	// send all the users as response
	err = json.NewEncoder(w).Encode(&models.UserList{Users: users})
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
}

func getEmailFromTokenHeader(r *http.Request) (string, error) {
	tkn := r.Header.Get("x-authentication-token")

	if tkn == "" {
		logrus.Debugf("no token in header")
		return "", errors.New("could not find token in x-authentication-token header")
	}
	return models.ParseToken(tkn)
}
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var userToAuth models.User
	err := json.NewDecoder(r.Body).Decode(&userToAuth)
	if err != nil {
		logrus.Errorf("Unable to decode the request body.  %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if userToAuth.Email == "" || userToAuth.Password == "" {
		http.Error(w, "No username or password provided", http.StatusBadRequest)
		return
	}

	userFromDb, err := database.GetUser(userToAuth.Email)
	if err != nil || userFromDb == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if !userFromDb.CheckPasswordHash(userToAuth.Password) {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}
	tkn, err := models.CreateToken(userToAuth.Email)
	if err != nil {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}
	err = json.NewEncoder(w).Encode(models.TokenResponse{Token: tkn})
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	email, err := getEmailFromTokenHeader(r)
	if err != nil || email == "" {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var user models.User
	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logrus.Debugf("Unable to decode the request body.  %v", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	// call update user to update the user
	updatedRows, err := database.UpdateUser(email, user.FirstName, user.LastName)
	if err != nil {
		logrus.Errorf("failed updating user.  %v", err)
		http.Error(w, "Invalid Request", http.StatusInternalServerError)
		return
	}
	logrus.Debugf("User updated successfully. Total rows/record affected %v", updatedRows)
}
