package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTopicConfigInfoRequest Request Object
type ListAlarmTopicConfigInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListAlarmTopicConfigInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTopicConfigInfoRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTopicConfigInfoRequest", string(data)}, " ")
}
