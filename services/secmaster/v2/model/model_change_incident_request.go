package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIncidentRequest Request Object
type ChangeIncidentRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
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
