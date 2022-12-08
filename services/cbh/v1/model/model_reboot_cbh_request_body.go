package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重启CBH实例请求对象
type RebootCbhRequestBody struct {

	// 实例的server id
	InstanceId string `json:"instance_id"`

	Reboot *RebootType `json:"reboot"`
}

func (o RebootCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootCbhRequestBody struct{}"
	}

	return strings.Join([]string{"RebootCbhRequestBody", string(data)}, " ")
}
