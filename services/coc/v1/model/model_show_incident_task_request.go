package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIncidentTaskRequest Request Object
type ShowIncidentTaskRequest struct {

	// 事件单ID
	IncidentId string `json:"incident_id"`
}

func (o ShowIncidentTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIncidentTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowIncidentTaskRequest", string(data)}, " ")
}
