package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDiagnosisReportRequest Request Object
type BatchDeleteDiagnosisReportRequest struct {

	// **参数解释**： 引擎。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteDiagnosisReportReq `json:"body,omitempty"`
}

func (o BatchDeleteDiagnosisReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDiagnosisReportRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDiagnosisReportRequest", string(data)}, " ")
}
