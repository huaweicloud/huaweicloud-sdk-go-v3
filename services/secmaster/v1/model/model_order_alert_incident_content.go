package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderAlertIncidentContent 事件内容
type OrderAlertIncidentContent struct {

	// 事件名称
	Title *string `json:"title,omitempty"`

	IncidentType *OrderAlertIncidentContentIncidentType `json:"incident_type,omitempty"`
}

func (o OrderAlertIncidentContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlertIncidentContent struct{}"
	}

	return strings.Join([]string{"OrderAlertIncidentContent", string(data)}, " ")
}
