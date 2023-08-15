package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorBrief 微服务引擎专享版规格及描述
type FlavorBrief struct {

	// 微服务引擎专享版规格
	Flavor *string `json:"flavor,omitempty"`

	// 微服务引擎专享版规格描述
	Description *string `json:"description,omitempty"`
}

func (o FlavorBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorBrief struct{}"
	}

	return strings.Join([]string{"FlavorBrief", string(data)}, " ")
}
