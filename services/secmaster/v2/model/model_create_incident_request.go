package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIncidentRequest Request Object
type CreateIncidentRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
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
