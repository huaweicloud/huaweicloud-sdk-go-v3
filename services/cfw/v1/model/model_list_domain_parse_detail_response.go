package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseDetailResponse Response Object
type ListDomainParseDetailResponse struct {

	// **参数解释**： 域名解析IP列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDomainParseDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDomainParseDetailResponse", string(data)}, " ")
}
