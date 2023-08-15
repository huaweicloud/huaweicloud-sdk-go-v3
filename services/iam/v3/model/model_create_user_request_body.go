package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserRequestBody
type CreateUserRequestBody struct {
	User *CreateUserOption `json:"user"`
}

func (o CreateUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUserRequestBody", string(data)}, " ")
}
