package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServicePodEventResponse struct {

	// **参数解释：** 事件发生次数。 **取值范围：** 不涉及。
	Count int32 `json:"count"`

	// **参数解释：** 首次发生时间。 **取值范围：** 不涉及。
	FirstTimestamp *sdktime.SdkTime `json:"first_timestamp"`

	// **参数解释：** 最近发生时间。 **取值范围：** 不涉及。
	LastTimestamp *sdktime.SdkTime `json:"last_timestamp"`

	// **参数解释：** 事件信息。 **取值范围：** 不涉及。
	Message string `json:"message"`

	// **参数解释：** 事件名称。 **取值范围：** 不涉及。
	Reason string `json:"reason"`

	// **参数解释：** 上报该事件的k8s组件名。 **取值范围：** 不涉及。
	ReportingComponent string `json:"reporting_component"`

	// **参数解释：** 事件类型。 **取值范围：** Normal/Warning。
	Type string `json:"type"`
}

func (o ServicePodEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServicePodEventResponse struct{}"
	}

	return strings.Join([]string{"ServicePodEventResponse", string(data)}, " ")
}
