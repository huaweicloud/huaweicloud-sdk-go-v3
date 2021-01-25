package model

import (
	"encoding/json"

	"strings"
)

// 执行复制请求体
type CheckpointReplicateReq struct {
	Replicate *CheckpointReplicateParam `json:"replicate"`
}

func (o CheckpointReplicateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckpointReplicateReq struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateReq", string(data)}, " ")
}
