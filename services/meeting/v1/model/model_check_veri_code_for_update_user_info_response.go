package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckVeriCodeForUpdateUserInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckVeriCodeForUpdateUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckVeriCodeForUpdateUserInfoResponse struct{}"
	}

	return strings.Join([]string{"CheckVeriCodeForUpdateUserInfoResponse", string(data)}, " ")
}
