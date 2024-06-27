package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResumeStarRocksDataReplication struct {

	// 同步任务名。字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线。
	TaskName string `json:"task_name"`
}

func (o ResumeStarRocksDataReplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeStarRocksDataReplication struct{}"
	}

	return strings.Join([]string{"ResumeStarRocksDataReplication", string(data)}, " ")
}
