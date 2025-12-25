package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderAlert 告警批量转事件body体
type OrderAlert struct {

	// 转事件的告警id列表
	Ids *[]string `json:"ids,omitempty"`

	IncidentContent *OrderAlertIncidentContent `json:"incident_content,omitempty"`
}

func (o OrderAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlert struct{}"
	}

	return strings.Join([]string{"OrderAlert", string(data)}, " ")
}
