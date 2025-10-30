package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeVulWhiteListRequestBody struct {

	// **参数解释**： 白名单规则类型 **约束限制**： 不涉及 **取值范围**： - all_host：白名单应用于全部主机 - specific_host：白名单应用于指定主机  **默认取值**： 不涉及
	RuleType string `json:"rule_type"`

	// **参数解释**： 白名单应用的主机ID列表，rule_type为specific_host时必填 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**： 白名单的描述信息 **约束限制**： 不涉及 **取值范围**： 字符长度0-1024 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o ChangeVulWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeVulWhiteListRequestBody", string(data)}, " ")
}
