package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceInfoForRisk 实例信息
type InstanceInfoForRisk struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o InstanceInfoForRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfoForRisk struct{}"
	}

	return strings.Join([]string{"InstanceInfoForRisk", string(data)}, " ")
}
