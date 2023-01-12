package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 事件类型
type OrderAlertIncidentContentIncidentType struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// Id value
	LayoutId *string `json:"layoutId,omitempty"`

	// Id value
	IncidentType *string `json:"incident_type,omitempty"`

	// Id value
	Category *string `json:"category,omitempty"`
}

func (o OrderAlertIncidentContentIncidentType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlertIncidentContentIncidentType struct{}"
	}

	return strings.Join([]string{"OrderAlertIncidentContentIncidentType", string(data)}, " ")
}
