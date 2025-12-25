package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeRequest Request Object
type DeleteNodeRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteNodeRequestBody `json:"body,omitempty"`
}

func (o DeleteNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodeRequest", string(data)}, " ")
}
