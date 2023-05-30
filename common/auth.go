package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"

	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath string = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey string
)

func initKeys() {
	var err error

	_, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	_, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}

}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
	// create a singer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set claims for JWT token
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expiring time for JWT token
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//  validate token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		//  verify the token
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case *jwt.ValidationErrorExpired: // JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Access Token",
					500,
				)
				return
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing the Access Token",
				500,
			)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		DisplayAppError(w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
