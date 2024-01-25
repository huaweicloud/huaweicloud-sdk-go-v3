package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Dubbo struct {

	// 服务名。
	Service *string `json:"service,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	// 分组。
	Group *string `json:"group,omitempty"`

	// dubbo方法列表。
	Methods *[]DubboMethod `json:"methods,omitempty"`
}

func (o Dubbo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dubbo struct{}"
	}

	return strings.Join([]string{"Dubbo", string(data)}, " ")
}
