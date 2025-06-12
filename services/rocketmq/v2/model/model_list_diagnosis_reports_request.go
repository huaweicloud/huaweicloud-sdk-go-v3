package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnosisReportsRequest Request Object
type ListDiagnosisReportsRequest struct {

	// **参数解释**： 引擎。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：登录RocketMQ控制台，在RocketMQ实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDiagnosisReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisReportsRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnosisReportsRequest", string(data)}, " ")
}
