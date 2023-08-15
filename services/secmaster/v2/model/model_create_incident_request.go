package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIncidentRequest Request Object
type CreateIncidentRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *CreateIncidentRequestBody `json:"body,omitempty"`
}

func (o CreateIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncidentRequest struct{}"
	}

	return strings.Join([]string{"CreateIncidentRequest", string(data)}, " ")
}
