package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateJobsReq 批量操作Job。
type BatchOperateJobsReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 50]。
	Items []string `json:"items"`
}

func (o BatchOperateJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateJobsReq struct{}"
	}

	return strings.Join([]string{"BatchOperateJobsReq", string(data)}, " ")
}
