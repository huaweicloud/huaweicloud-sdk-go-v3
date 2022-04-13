package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量结束与删除在线迁移任务请求体
type BatchDeleteJobReq struct {
	// 批量结束与删除任务请求列表

	Jobs []DeleteJobReq `json:"jobs"`
}

func (o BatchDeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobReq", string(data)}, " ")
}
