package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderAlert 转事件
type OrderAlert struct {

	// 转事件的ID列表
	Ids *[]string `json:"ids,omitempty"`

	// 事件id
	IncidentId *string `json:"incident_id,omitempty"`

	EventContent *OrderAlertEventContent `json:"event_content,omitempty"`

	IncidentContent *OrderAlertIncidentContent `json:"incident_content,omitempty"`

	// 标记为证据
	MarkedEvidence *bool `json:"marked_evidence,omitempty"`
}

func (o OrderAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlert struct{}"
	}

	return strings.Join([]string{"OrderAlert", string(data)}, " ")
}
