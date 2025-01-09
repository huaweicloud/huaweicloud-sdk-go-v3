package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScheduleTaskReq 批量删除定时任务。
type BatchDeleteScheduleTaskReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]
	Items []string `json:"items"`
}

func (o BatchDeleteScheduleTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScheduleTaskReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteScheduleTaskReq", string(data)}, " ")
}
