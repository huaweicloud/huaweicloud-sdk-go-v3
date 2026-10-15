package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetResultDto struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 是否设置成功（true-成功 false-失败）。
	Success bool `json:"success"`

	// 失败时的错误信息，成功时为null。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchSetResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetResultDto struct{}"
	}

	return strings.Join([]string{"BatchSetResultDto", string(data)}, " ")
}
