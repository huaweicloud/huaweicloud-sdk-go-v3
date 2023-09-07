package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobActionReq 结束单个任务请求体。
type StopJobActionReq struct {

	// 强制结束任务时取值true，默认false。
	IsForceStop *bool `json:"is_force_stop,omitempty"`
}

func (o StopJobActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobActionReq struct{}"
	}

	return strings.Join([]string{"StopJobActionReq", string(data)}, " ")
}
