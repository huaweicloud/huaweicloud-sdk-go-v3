package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EcsSpecificationBean struct {

	// 可用区集合
	Azs []string `json:"azs"`

	// ID
	Id string `json:"id"`

	// 等级
	Level string `json:"level"`

	// 名称
	Name string `json:"name"`

	// 代理
	Proxy int32 `json:"proxy"`

	// 内存
	Ram int32 `json:"ram"`

	// CPU
	Vcpus int32 `json:"vcpus"`
}

func (o EcsSpecificationBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EcsSpecificationBean struct{}"
	}

	return strings.Join([]string{"EcsSpecificationBean", string(data)}, " ")
}
