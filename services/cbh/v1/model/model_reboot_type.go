package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootType 云堡垒机实例重启方式。
type RebootType struct {

	// 重启方式，不区分大小写。 - SOFT  普通重启，关闭虚拟机服务 - HARD  强制重启，重启虚拟机
	Type string `json:"type"`
}

func (o RebootType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootType struct{}"
	}

	return strings.Join([]string{"RebootType", string(data)}, " ")
}
