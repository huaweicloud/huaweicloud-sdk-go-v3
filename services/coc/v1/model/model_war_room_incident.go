package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WarRoomIncident WarRoom存储事件有关信息
type WarRoomIncident struct {

	// 主键
	Id string `json:"id"`

	// 事件id
	IncidentId *string `json:"incident_id,omitempty"`

	// 是否变更事件
	IsChangeEvent *bool `json:"is_change_event,omitempty"`

	// 事件级别
	FailureLevel *string `json:"failure_level,omitempty"`

	// 事件单号链接
	IncidentUrl *string `json:"incident_url,omitempty"`
}

func (o WarRoomIncident) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarRoomIncident struct{}"
	}

	return strings.Join([]string{"WarRoomIncident", string(data)}, " ")
}
