package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量查询同步进度的请求体
type BatchQueryProgressReq struct {

	// 批量查询进度任务ID请求列表
	Jobs []string `json:"jobs" xml:"jobs"`
}

func (o BatchQueryProgressReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryProgressReq struct{}"
	}

	return strings.Join([]string{"BatchQueryProgressReq", string(data)}, " ")
}
