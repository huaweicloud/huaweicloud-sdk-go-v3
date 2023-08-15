package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordResponse Response Object
type ChangePasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordResponse struct{}"
	}

	return strings.Join([]string{"ChangePasswordResponse", string(data)}, " ")
}
