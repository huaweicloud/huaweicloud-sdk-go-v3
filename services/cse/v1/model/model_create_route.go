package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoute 路由规则。
type CreateRoute struct {

	// 规则名称。
	Name *string `json:"name,omitempty"`

	// 权重值。
	Weight *int32 `json:"weight,omitempty"`

	Tags *CreateRouteTags `json:"tags,omitempty"`
}

func (o CreateRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoute struct{}"
	}

	return strings.Join([]string{"CreateRoute", string(data)}, " ")
}
