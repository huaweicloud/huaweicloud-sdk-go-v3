package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartJobReq 批量启动任务请求体。
type BatchStartJobReq struct {

	// 批量启动任务请求列表。
	Jobs []StartInfo `json:"jobs"`
}

func (o BatchStartJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobReq struct{}"
	}

	return strings.Join([]string{"BatchStartJobReq", string(data)}, " ")
}
