package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncidentRequest Request Object
type DeleteIncidentRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteIncidentRequestBody `json:"body,omitempty"`
}

func (o DeleteIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncidentRequest struct{}"
	}

	return strings.Join([]string{"DeleteIncidentRequest", string(data)}, " ")
}
