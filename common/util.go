package common

import (
	"context"
	"encoding/json"
	"github.com/stkr89/go-k8s-cognito-crud-microservice-template/types"
)

func GetUserFromContext(ctx context.Context) (*types.User, error) {
	errMsg := "user not found"

	userBytes := ctx.Value("user")
	if userBytes == nil {
		return nil, NewError(UserNotFoundInContext, errMsg)
	}

	var user types.User
	err := json.Unmarshal(userBytes.([]byte), &user)
	if err != nil {
		return nil, NewError(UserNotFoundInContext, errMsg)
	}
	return &user, nil
}
