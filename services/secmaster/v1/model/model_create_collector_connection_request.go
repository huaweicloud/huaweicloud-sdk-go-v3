package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorConnectionRequest Request Object
type CreateCollectorConnectionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateConnectionDto `json:"body,omitempty"`
}

func (o CreateCollectorConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectorConnectionRequest", string(data)}, " ")
}
