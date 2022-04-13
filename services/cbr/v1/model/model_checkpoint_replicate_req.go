package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行复制请求体
type CheckpointReplicateReq struct {
	Replicate *CheckpointReplicateParam `json:"replicate"`
}

func (o CheckpointReplicateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateReq struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateReq", string(data)}, " ")
}
