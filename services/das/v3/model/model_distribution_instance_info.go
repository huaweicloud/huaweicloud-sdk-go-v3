package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DistributionInstanceInfo DistributionInstanceInfo对象
type DistributionInstanceInfo struct {

	// 实例状态
	Status *string `json:"status,omitempty"`

	// 实例数量
	Num *int32 `json:"num,omitempty"`
}

func (o DistributionInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistributionInstanceInfo struct{}"
	}

	return strings.Join([]string{"DistributionInstanceInfo", string(data)}, " ")
}
