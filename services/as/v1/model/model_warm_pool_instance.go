package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WarmPoolInstance struct {

	// 暖池实例ID
	Id *string `json:"id,omitempty"`

	// 对应的虚拟机ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 暖池实例的名称
	Name *string `json:"name,omitempty"`

	// 暖池实例的状态
	Status *string `json:"status,omitempty"`
}

func (o WarmPoolInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarmPoolInstance struct{}"
	}

	return strings.Join([]string{"WarmPoolInstance", string(data)}, " ")
}
