package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchTransferTask struct {

	// 转储任务ID
	Id string `json:"id"`
}

func (o BatchTransferTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTransferTask struct{}"
	}

	return strings.Join([]string{"BatchTransferTask", string(data)}, " ")
}
