package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDnsServersRequestBodyDnsServer struct {

	// DNS服务器IP
	ServerIp *string `json:"server_ip,omitempty"`

	// 是否是用户自定义的dns服务器，0否 1是
	IsCustomized *int32 `json:"is_customized,omitempty"`

	// 是否应用，0否 1是
	IsApplied *int32 `json:"is_applied,omitempty"`
}

func (o UpdateDnsServersRequestBodyDnsServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequestBodyDnsServer struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequestBodyDnsServer", string(data)}, " ")
}
