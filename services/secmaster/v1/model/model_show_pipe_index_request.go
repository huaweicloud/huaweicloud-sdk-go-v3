package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipeIndexRequest Request Object
type ShowPipeIndexRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`
}

func (o ShowPipeIndexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipeIndexRequest struct{}"
	}

	return strings.Join([]string{"ShowPipeIndexRequest", string(data)}, " ")
}
