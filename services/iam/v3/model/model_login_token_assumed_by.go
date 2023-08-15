package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginTokenAssumedBy
type LoginTokenAssumedBy struct {
	User *LoginTokenUser `json:"user,omitempty"`
}

func (o LoginTokenAssumedBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginTokenAssumedBy struct{}"
	}

	return strings.Join([]string{"LoginTokenAssumedBy", string(data)}, " ")
}
