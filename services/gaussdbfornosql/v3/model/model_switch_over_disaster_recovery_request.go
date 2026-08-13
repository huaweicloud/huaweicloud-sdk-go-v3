package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOverDisasterRecoveryRequest Request Object
type SwitchOverDisasterRecoveryRequest struct {

	// **参数解释：** 源实例ID或容灾实例ID。实例ID可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`
}

func (o SwitchOverDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOverDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"SwitchOverDisasterRecoveryRequest", string(data)}, " ")
}
