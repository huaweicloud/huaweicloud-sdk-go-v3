package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamedContext 上下文
type NamedContext struct {

	// 上下文的名称
	Name *string `json:"name,omitempty"`

	Context *Context `json:"context,omitempty"`
}

func (o NamedContext) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamedContext struct{}"
	}

	return strings.Join([]string{"NamedContext", string(data)}, " ")
}
