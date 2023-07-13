package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobActionReq 操作单个任务请求体，支持测试连接、预检查、启动、暂停、续传、重置、对比、结束等操作。
type JobActionReq struct {
	Job *ActionReq `json:"job"`
}

func (o JobActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobActionReq struct{}"
	}

	return strings.Join([]string{"JobActionReq", string(data)}, " ")
}
