package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseIpListResponse **参数解释**： 查询域名解析的ip列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type ParseIpListResponse struct {

	// **参数解释**： 超过限制的ip列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ExcessIp *[]string `json:"excess_ip,omitempty"`

	// **参数解释**： 成功解析的ip列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ParsedSuccessIp *[]string `json:"parsed_success_ip,omitempty"`
}

func (o ParseIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseIpListResponse struct{}"
	}

	return strings.Join([]string{"ParseIpListResponse", string(data)}, " ")
}
