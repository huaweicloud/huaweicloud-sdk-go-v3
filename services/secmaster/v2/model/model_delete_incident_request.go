package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncidentRequest Request Object
type DeleteIncidentRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteIncident `json:"body,omitempty"`
}

func (o DeleteIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncidentRequest struct{}"
	}

	return strings.Join([]string{"DeleteIncidentRequest", string(data)}, " ")
}
