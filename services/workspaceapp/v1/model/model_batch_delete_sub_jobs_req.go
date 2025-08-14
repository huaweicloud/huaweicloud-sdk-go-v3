package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubJobsReq 批量删除子任务请求。
type BatchDeleteSubJobsReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]。
	Items []string `json:"items"`
}

func (o BatchDeleteSubJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubJobsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubJobsReq", string(data)}, " ")
}
