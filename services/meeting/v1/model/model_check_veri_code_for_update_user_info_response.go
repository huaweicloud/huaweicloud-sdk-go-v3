package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckVeriCodeForUpdateUserInfoResponse Response Object
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
