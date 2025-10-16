package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatusInstanceItemInstanceList struct {

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 采集值
	Value *int32 `json:"value,omitempty"`
}

func (o ShowStatusInstanceItemInstanceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusInstanceItemInstanceList struct{}"
	}

	return strings.Join([]string{"ShowStatusInstanceItemInstanceList", string(data)}, " ")
}
