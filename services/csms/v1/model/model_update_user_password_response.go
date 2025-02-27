package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserPasswordResponse Response Object
type UpdateUserPasswordResponse struct {

	// 用户id。
	UserId         *string `json:"user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserPasswordResponse", string(data)}, " ")
}
