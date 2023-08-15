package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDnsServersRequestBody struct {

	// DNS服务器
	DnsServer *[]UpdateDnsServersRequestBodyDnsServer `json:"dns_server,omitempty"`
}

func (o UpdateDnsServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequestBody", string(data)}, " ")
}
