package opaque_token

import (
	"github.com/go-playground/validator/v10"
)

type OpaqueToken struct {
	Token string `validate:"uuid5"`
	Jwt   string `validate:"omitempty,jwt"`
}

func New(token string) (*OpaqueToken, error) {
	t := OpaqueToken{
		Token: token,
	}

	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func NewWithJwt(token string, jwt string) (*OpaqueToken, error) {
	t := OpaqueToken{
		Token: token,
		Jwt:   jwt,
	}

	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (op *OpaqueToken) SetJwt(jwt string) {
	op.Jwt = jwt
}
