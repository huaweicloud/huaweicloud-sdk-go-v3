package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单条日志记录
type CheckpointRsp struct {

	// 数据名称
	Source *string `json:"source,omitempty"`

	// 日志时间戳
	Timestamp *string `json:"timestamp,omitempty"`

	// 执行信息
	Message *string `json:"message,omitempty"`
}

func (o CheckpointRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointRsp struct{}"
	}

	return strings.Join([]string{"CheckpointRsp", string(data)}, " ")
}
