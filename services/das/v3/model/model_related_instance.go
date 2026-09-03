package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelatedInstance 相关实例信息
type RelatedInstance struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例类型
	Type *string `json:"type,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`
}

func (o RelatedInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedInstance struct{}"
	}

	return strings.Join([]string{"RelatedInstance", string(data)}, " ")
}
