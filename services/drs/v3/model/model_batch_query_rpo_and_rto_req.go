package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryRpoAndRtoReq 批量查询RPO和RTO的请求体
type BatchQueryRpoAndRtoReq struct {

	// 批量查询RPO和RTO的任务详情ID请求列表
	Jobs []string `json:"jobs"`
}

func (o BatchQueryRpoAndRtoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryRpoAndRtoReq struct{}"
	}

	return strings.Join([]string{"BatchQueryRpoAndRtoReq", string(data)}, " ")
}
