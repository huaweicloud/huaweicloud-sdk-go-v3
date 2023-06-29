package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobReq 批量删除任务请求体。
type BatchDeleteJobReq struct {

	// 批量删除任务请求体。
	Jobs []string `json:"jobs"`
}

func (o BatchDeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobReq", string(data)}, " ")
}
