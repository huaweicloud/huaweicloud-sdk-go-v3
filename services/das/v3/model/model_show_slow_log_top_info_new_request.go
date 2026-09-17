package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogTopInfoNewRequest Request Object
type ShowSlowLogTopInfoNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`
}

func (o ShowSlowLogTopInfoNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogTopInfoNewRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogTopInfoNewRequest", string(data)}, " ")
}
