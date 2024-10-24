package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanupModelRequest Request Object
type CleanupModelRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 推理模型ID
	ModelId string `json:"model_id"`
}

func (o CleanupModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanupModelRequest struct{}"
	}

	return strings.Join([]string{"CleanupModelRequest", string(data)}, " ")
}
