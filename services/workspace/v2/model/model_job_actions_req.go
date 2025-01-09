package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobActionsReq 重试任务请求。
type JobActionsReq struct {

	// 操作类型。retry代表重试。
	OpType string `json:"op_type"`
}

func (o JobActionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobActionsReq struct{}"
	}

	return strings.Join([]string{"JobActionsReq", string(data)}, " ")
}
