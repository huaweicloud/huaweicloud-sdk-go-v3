package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupRelationConfiguration 分组的关联配置信息，比如对应的APM的配置信息。
type GroupRelationConfiguration struct {

	// 配置类型。
	Type *string `json:"type,omitempty"`

	// 配置参数。
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (o GroupRelationConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupRelationConfiguration struct{}"
	}

	return strings.Join([]string{"GroupRelationConfiguration", string(data)}, " ")
}
