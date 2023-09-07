package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopJobActionReq 批量结束任务请求体。
type BatchStopJobActionReq struct {

	// 批量结束任务请求体。
	Jobs []StopJobActionInfo `json:"jobs"`
}

func (o BatchStopJobActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopJobActionReq struct{}"
	}

	return strings.Join([]string{"BatchStopJobActionReq", string(data)}, " ")
}
