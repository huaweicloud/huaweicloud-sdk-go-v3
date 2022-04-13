package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量创建实时迁移任务请求体
type BatchCreateJobReq struct {
	// 创建任务请求体

	Jobs []CreateJobReq `json:"jobs"`
}

func (o BatchCreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobReq struct{}"
	}

	return strings.Join([]string{"BatchCreateJobReq", string(data)}, " ")
}
