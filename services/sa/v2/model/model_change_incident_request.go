package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIncidentRequest Request Object
type ChangeIncidentRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// 事件ID
	IncidentId string `json:"incident_id"`

	Body *ChangeIncidentRequestBody `json:"body,omitempty"`
}

func (o ChangeIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIncidentRequest struct{}"
	}

	return strings.Join([]string{"ChangeIncidentRequest", string(data)}, " ")
}
