package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderAlertIncidentContentIncidentType 事件类型
type OrderAlertIncidentContentIncidentType struct {

	// 事件类型id
	Id *string `json:"id,omitempty"`

	// 事件类型父类
	Category *string `json:"category,omitempty"`

	// 事件类型子类
	IncidentType *string `json:"incident_type,omitempty"`
}

func (o OrderAlertIncidentContentIncidentType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlertIncidentContentIncidentType struct{}"
	}

	return strings.Join([]string{"OrderAlertIncidentContentIncidentType", string(data)}, " ")
}
