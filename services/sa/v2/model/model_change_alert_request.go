package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeAlertRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// 告警ID
	AlertId string `json:"alert_id"`

	Body *ChangeAlertRequestBody `json:"body,omitempty"`
}

func (o ChangeAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlertRequest struct{}"
	}

	return strings.Join([]string{"ChangeAlertRequest", string(data)}, " ")
}
