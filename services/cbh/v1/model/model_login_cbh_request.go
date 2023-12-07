package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginCbhRequest Request Object
type LoginCbhRequest struct {
	Body *LoginCbhRequestBody `json:"body,omitempty"`
}

func (o LoginCbhRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginCbhRequest struct{}"
	}

	return strings.Join([]string{"LoginCbhRequest", string(data)}, " ")
}
