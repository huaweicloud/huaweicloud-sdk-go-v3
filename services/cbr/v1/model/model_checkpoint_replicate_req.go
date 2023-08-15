package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckpointReplicateReq 执行复制请求体
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
