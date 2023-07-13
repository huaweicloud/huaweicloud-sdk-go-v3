package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootCbhRequestBody 重启云堡垒机实例请求对象。
type RebootCbhRequestBody struct {

	// 云堡垒机实例ID，使用UUID格式。
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
