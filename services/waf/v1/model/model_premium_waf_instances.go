package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PremiumWafInstances 租户引擎实例信息
type PremiumWafInstances struct {

	// 引擎实例id
	Id *string `json:"id,omitempty"`

	// 引擎实例名
	Name *string `json:"name,omitempty"`

	// 引擎实例是否已接入，false：未接入；true：已接入
	Accessed *bool `json:"accessed,omitempty"`
}

func (o PremiumWafInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafInstances struct{}"
	}

	return strings.Join([]string{"PremiumWafInstances", string(data)}, " ")
}
