package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginWhiteIpResponse Response Object
type ModifyLoginWhiteIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyLoginWhiteIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginWhiteIpResponse struct{}"
	}

	return strings.Join([]string{"ModifyLoginWhiteIpResponse", string(data)}, " ")
}
