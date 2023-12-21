package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootCbhRequestBody 重启云堡垒机实例请求对象。
type RebootCbhRequestBody struct {
	ServerId *interface{} `json:"server_id"`

	// 重启方式，不区分大小写。 - SOFT：普通重启，关闭虚拟机服务 - HARD：强制重启，重启虚拟机
	RebootType string `json:"reboot_type"`
}

func (o RebootCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootCbhRequestBody struct{}"
	}

	return strings.Join([]string{"RebootCbhRequestBody", string(data)}, " ")
}
