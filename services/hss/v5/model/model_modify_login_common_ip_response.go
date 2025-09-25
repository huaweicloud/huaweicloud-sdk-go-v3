package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginCommonIpResponse Response Object
type ModifyLoginCommonIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyLoginCommonIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginCommonIpResponse struct{}"
	}

	return strings.Join([]string{"ModifyLoginCommonIpResponse", string(data)}, " ")
}
