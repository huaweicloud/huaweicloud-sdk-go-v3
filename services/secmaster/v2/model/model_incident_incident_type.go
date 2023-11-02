package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IncidentIncidentType 事件分类，详细定义参考《告警事件类型定义》
type IncidentIncidentType struct {

	// 类别
	Category *string `json:"category,omitempty"`

	// 事件类型
	IncidentType *string `json:"incident_type,omitempty"`
}

func (o IncidentIncidentType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentIncidentType struct{}"
	}

	return strings.Join([]string{"IncidentIncidentType", string(data)}, " ")
}
