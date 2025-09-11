package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSensitiveRulesRequest Request Object
type DeleteSensitiveRulesRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**： 隐私规则ID。可通过查询隐私数据脱敏规则列表ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询隐私数据脱敏规则列表值为准，字符长度16-64。 **默认取值**： 不涉及
	Id string `json:"id"`
}

func (o DeleteSensitiveRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSensitiveRulesRequest struct{}"
	}

	return strings.Join([]string{"DeleteSensitiveRulesRequest", string(data)}, " ")
}
