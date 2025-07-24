package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceIpResponse Response Object
type ModifyInstanceIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyInstanceIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceIpResponse struct{}"
	}

	return strings.Join([]string{"ModifyInstanceIpResponse", string(data)}, " ")
}
