package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDemand 资源需求量配置
type ResourceDemand struct {

	// 资源规格，从规格列表查询获取。
	SpecCode string `json:"spec_code"`

	// 最小数
	Min int32 `json:"min"`

	// 最大数
	Max int32 `json:"max"`
}

func (o ResourceDemand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDemand struct{}"
	}

	return strings.Join([]string{"ResourceDemand", string(data)}, " ")
}
