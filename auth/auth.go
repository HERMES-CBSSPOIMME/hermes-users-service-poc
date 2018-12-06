package auth

import (
	errors "errors"
	time "time"

	jwt "github.com/dgrijalva/jwt-go"

	models "hermes-users-service/models"
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
