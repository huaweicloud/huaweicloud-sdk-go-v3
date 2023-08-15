package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArchitectureStatistic struct {

	// 子指标
	Children *[]ArchitectureStatistic `json:"children,omitempty"`

	// 子指标个数
	Count *int32 `json:"count,omitempty"`

	// guid
	Guid *string `json:"guid,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o ArchitectureStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchitectureStatistic struct{}"
	}

	return strings.Join([]string{"ArchitectureStatistic", string(data)}, " ")
}
