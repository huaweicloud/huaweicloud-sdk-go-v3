package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRequest Request Object
type CreateAlertRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAlertRequestBody `json:"body,omitempty"`
}

func (o CreateAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRequest struct{}"
	}

	return strings.Join([]string{"CreateAlertRequest", string(data)}, " ")
}
