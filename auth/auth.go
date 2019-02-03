package auth

import (
	errors "errors"
	fmt "fmt"
	time "time"

	jwt "github.com/dgrijalva/jwt-go"

	models "wave-users-service-poc/models"
)

var (
	SigningKey = []byte("ThisisAnExampleSigningKey!PleaseDontUseThatInProduction!")
)

// TODO see if i keep this struct
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

type TokenWrap struct {
	SignedToken string `json:"jwt"`
}

type UidWrap struct {
	UserID string `json:"userID"`
}

func (creds *Credentials) Verify(env *models.Env) (string, error) {

	user, err := env.MongoDB.GetUserByUsername(creds.Username)

	if err != nil {
		return "", err
	}

	if user.Password != creds.Password {
		return "", errors.New("Bad Login")
	} else {
		return user.Uid, nil
	}

}

func NewCustomClaim(userID string) *CustomClaims {
	return &CustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "HermesDemo",
		},
	}
}

func CreateToken(userID string) (*TokenWrap, error) {
	claims := NewCustomClaim(userID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SigningKey)

	if err != nil {
		return nil, err
	}

	tw := TokenWrap{
		SignedToken: signedToken,
	}

	return &tw, nil
}

func ValidateToken(signedToken string) (*UidWrap, error) {
	var uw UidWrap

	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SigningKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		uw.UserID = claims.Uid
	} else {
		return nil, err
	}
	return &uw, nil
}
