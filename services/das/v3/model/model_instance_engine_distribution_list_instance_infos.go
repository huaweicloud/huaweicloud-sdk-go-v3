package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceEngineDistributionListInstanceInfos struct {

	// 实例状态
	Status *string `json:"status,omitempty"`

	// 实例数量
	Num *int32 `json:"num,omitempty"`
}

func (o InstanceEngineDistributionListInstanceInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceEngineDistributionListInstanceInfos struct{}"
	}

	return strings.Join([]string{"InstanceEngineDistributionListInstanceInfos", string(data)}, " ")
}
