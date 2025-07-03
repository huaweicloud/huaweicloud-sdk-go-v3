package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleIncidentRequest Request Object
type HandleIncidentRequest struct {

	// 事件单ID
	IncidentId string `json:"incident_id"`

	Body *ExecuteActionParamsV2 `json:"body,omitempty"`
}

func (o HandleIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleIncidentRequest struct{}"
	}

	return strings.Join([]string{"HandleIncidentRequest", string(data)}, " ")
}
