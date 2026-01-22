package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseIpResponse Response Object
type ListDomainParseIpResponse struct {

	// **参数解释**： 域名解析数据，包括域名IP列表 **取值范围**： 不涉及
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
