package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDef struct {

	// 资源规格编码,从查询规格列表ListSpecs接口获取
	SpecCode string `json:"spec_code"`

	// 预热资源量
	WarmUpNum int32 `json:"warm_up_num"`

	// 最大资源量，不填默认为预热资源量，即不使用弹性资源
	MaxNum *int32 `json:"max_num,omitempty"`
}

func (o ResourceDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDef struct{}"
	}

	return strings.Join([]string{"ResourceDef", string(data)}, " ")
}
