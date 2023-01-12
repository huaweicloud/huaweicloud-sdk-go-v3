package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 事件类型
type CreateIncidentIncidentType struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// Id value
	LayoutId *string `json:"layoutId,omitempty"`

	// Id value
	IncidentType *string `json:"incident_type,omitempty"`

	// Id value
	Category *string `json:"category,omitempty"`
}

func (o CreateIncidentIncidentType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncidentIncidentType struct{}"
	}

	return strings.Join([]string{"CreateIncidentIncidentType", string(data)}, " ")
}
