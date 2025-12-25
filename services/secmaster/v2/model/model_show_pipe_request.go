package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipeRequest Request Object
type ShowPipeRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`
}

func (o ShowPipeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipeRequest struct{}"
	}

	return strings.Join([]string{"ShowPipeRequest", string(data)}, " ")
}
