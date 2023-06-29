package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPwdByAdminResponse Response Object
type ResetPwdByAdminResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdByAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdByAdminResponse struct{}"
	}

	return strings.Join([]string{"ResetPwdByAdminResponse", string(data)}, " ")
}
