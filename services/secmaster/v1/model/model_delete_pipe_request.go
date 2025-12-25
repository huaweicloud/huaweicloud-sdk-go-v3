package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipeRequest Request Object
type DeletePipeRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`
}

func (o DeletePipeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipeRequest struct{}"
	}

	return strings.Join([]string{"DeletePipeRequest", string(data)}, " ")
}
