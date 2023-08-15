package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Instances struct {

	// 缓存实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 缓存实例名称。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o Instances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instances struct{}"
	}

	return strings.Join([]string{"Instances", string(data)}, " ")
}
