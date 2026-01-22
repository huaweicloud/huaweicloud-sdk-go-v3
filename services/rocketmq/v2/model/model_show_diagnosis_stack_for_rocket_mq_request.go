package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisStackForRocketMqRequest Request Object
type ShowDiagnosisStackForRocketMqRequest struct {

	// **参数解释**： 堆ID。 从查询实例诊断报告接口获取。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StackId string `json:"stack_id"`
}

func (o ShowDiagnosisStackForRocketMqRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisStackForRocketMqRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisStackForRocketMqRequest", string(data)}, " ")
}
