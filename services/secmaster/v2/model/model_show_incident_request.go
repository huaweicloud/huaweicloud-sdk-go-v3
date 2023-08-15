package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIncidentRequest Request Object
type ShowIncidentRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// 事件ID
	IncidentId string `json:"incident_id"`
}

func (o ShowIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIncidentRequest struct{}"
	}

	return strings.Join([]string{"ShowIncidentRequest", string(data)}, " ")
}
