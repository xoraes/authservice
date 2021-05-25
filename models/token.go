package models

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)
var JWTKEY = os.Getenv("JWT_KEY")

type TokenResponse struct {
	Token string `json:"token,omitempty"`
}

type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(email string) (string, error) {
	out := os.Getenv("JWT_TIMEOUT_MINUTES")
	to, err := strconv.Atoi(out)

	if err != nil {
		to = 5
	}
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(time.Duration(to) * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	//TODO fix this
	tokenString, err := token.SignedString([]byte(JWTKEY))
	return tokenString, err
}

func ParseToken(token string) (string, error) {
	// Initialize a new instance of `Claims`
	claims := &claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	//TODO fix this key
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKEY), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", appErr(err.Error(), UNAUTHORIZED)
		}
		return "", appErr(err.Error(), BADREQUEST)
	}
	if !tkn.Valid {
		return "", appErr("Invalid Token", UNAUTHORIZED)
	}
	return claims.Email, nil
}
