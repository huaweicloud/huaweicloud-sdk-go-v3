package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorBrief 微服务引擎规格及描述
type FlavorBrief struct {

	// 微服务引擎规格
	Flavor *string `json:"flavor,omitempty"`

	// 微服务引擎规格描述
	Description *string `json:"description,omitempty"`

	Spec *EngineSpec `json:"spec,omitempty"`
}

func (o FlavorBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorBrief struct{}"
	}

	return strings.Join([]string{"FlavorBrief", string(data)}, " ")
}
