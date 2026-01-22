package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDiagnosisRecordsForRocketMqRequest Request Object
type BatchDeleteDiagnosisRecordsForRocketMqRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteDiagnosisReportReq `json:"body,omitempty"`
}

func (o BatchDeleteDiagnosisRecordsForRocketMqRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDiagnosisRecordsForRocketMqRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDiagnosisRecordsForRocketMqRequest", string(data)}, " ")
}
