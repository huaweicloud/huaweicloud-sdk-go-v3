package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseIpResponse Response Object
type ListDomainParseIpResponse struct {

	// 域名解析数据
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDomainParseIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseIpResponse struct{}"
	}

	return strings.Join([]string{"ListDomainParseIpResponse", string(data)}, " ")
}
