package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRequest Request Object
type ShowAlertRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// 告警ID
	AlertId string `json:"alert_id"`
}

func (o ShowAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertRequest", string(data)}, " ")
}
