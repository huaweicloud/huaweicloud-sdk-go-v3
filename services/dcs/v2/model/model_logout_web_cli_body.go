package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogoutWebCliBody struct {

	// 客户端ID
	ClientId *string `json:"client_id,omitempty"`
}

func (o LogoutWebCliBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoutWebCliBody struct{}"
	}

	return strings.Join([]string{"LogoutWebCliBody", string(data)}, " ")
}
