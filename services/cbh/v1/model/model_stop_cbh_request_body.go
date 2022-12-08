package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 关闭CBH实例请求对象
type StopCbhRequestBody struct {

	// 实例的server id
	InstanceId string `json:"instance_id"`

	Reboot *RebootType `json:"reboot,omitempty"`
}

func (o StopCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCbhRequestBody struct{}"
	}

	return strings.Join([]string{"StopCbhRequestBody", string(data)}, " ")
}
