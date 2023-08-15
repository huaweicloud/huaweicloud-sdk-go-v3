package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRequest Request Object
type DeleteAlertRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteAlert `json:"body,omitempty"`
}

func (o DeleteAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlertRequest", string(data)}, " ")
}
