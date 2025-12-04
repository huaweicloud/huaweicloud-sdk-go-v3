package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoCesAlarmResponse Response Object
type ShowAutoCesAlarmResponse struct {

	// **参数解释**：  自动告警列表  **约束限制**：  不涉及。
	Entities *[]AutoCesAlarmInfo `json:"entities,omitempty"`

	// **参数解释**：  总数量。  **约束限制**：  不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAutoCesAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoCesAlarmResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoCesAlarmResponse", string(data)}, " ")
}
