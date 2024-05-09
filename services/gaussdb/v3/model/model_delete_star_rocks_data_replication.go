package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRocksDataReplication 删除StarRocks数据同步任务请求体。
type DeleteStarRocksDataReplication struct {

	// 数据同步任务名。字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线_。
	TaskName *string `json:"task_name,omitempty"`
}

func (o DeleteStarRocksDataReplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRocksDataReplication struct{}"
	}

	return strings.Join([]string{"DeleteStarRocksDataReplication", string(data)}, " ")
}
