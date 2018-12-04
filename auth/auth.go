package auth

import (
	errors "errors"

	models "hermes-users-service/models"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (creds *Credential) Verify(env *models.Env) (string, error) {

	user, err := env.MongoDB.GetUserByUsername(creds.Username)

	if err != nil {
		return "", err
	}

	if user.Password == creds.Password {
		return user.Uid, nil
	} else {
		return "", errors.New("Bad Login")
	}

}
