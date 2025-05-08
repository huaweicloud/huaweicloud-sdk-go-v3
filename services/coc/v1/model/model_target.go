package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Target struct {

	// 速率模式执行类型,InstanceValues
	Key string `json:"key"`

	// 速率模式执行元素，目前只支持cmdb资源
	Values []interface{} `json:"values"`
}

func (o Target) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Target struct{}"
	}

	return strings.Join([]string{"Target", string(data)}, " ")
}
