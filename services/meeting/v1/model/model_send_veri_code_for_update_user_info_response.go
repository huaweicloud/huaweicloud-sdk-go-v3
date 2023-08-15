package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendVeriCodeForUpdateUserInfoResponse Response Object
type SendVeriCodeForUpdateUserInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendVeriCodeForUpdateUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVeriCodeForUpdateUserInfoResponse struct{}"
	}

	return strings.Join([]string{"SendVeriCodeForUpdateUserInfoResponse", string(data)}, " ")
}
