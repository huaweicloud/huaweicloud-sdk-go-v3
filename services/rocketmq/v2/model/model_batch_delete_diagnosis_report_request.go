package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDiagnosisReportRequest Request Object
type BatchDeleteDiagnosisReportRequest struct {

	// **参数解释**： 引擎。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：登录RocketMQ控制台，在RocketMQ实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
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
