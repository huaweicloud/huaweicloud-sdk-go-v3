package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchJobActionReq 批量操作任务请求体，支持测试连接、预检查、启动、暂停、续传、重置、对比、结束等操作。
type BatchJobActionReq struct {

	// 批量操作任务请求体。
	Jobs []ActionReq `json:"jobs"`
}

func (o BatchJobActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchJobActionReq struct{}"
	}

	return strings.Join([]string{"BatchJobActionReq", string(data)}, " ")
}
