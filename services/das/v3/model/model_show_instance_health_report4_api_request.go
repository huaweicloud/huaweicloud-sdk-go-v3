package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceHealthReport4ApiRequest Request Object
type ShowInstanceHealthReport4ApiRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 日报任务ID
	TaskId string `json:"task_id"`
}

func (o ShowInstanceHealthReport4ApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceHealthReport4ApiRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceHealthReport4ApiRequest", string(data)}, " ")
}
