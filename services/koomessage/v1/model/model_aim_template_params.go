package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimTemplateParams 模板参数。
type AimTemplateParams struct {

	// 参数类型。 - string：文本 - integer：数字
	Type *string `json:"type,omitempty"`

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 动态参数是否长度限制。 - false：不可设置  - true：可设置
	HasLength *bool `json:"has_length,omitempty"`

	// 固定长度。
	FixLength *int32 `json:"fix_length,omitempty"`

	// 长度限制。
	LengthRestrict *bool `json:"length_restrict,omitempty"`

	// 最小长度。
	MinLength *int32 `json:"min_length,omitempty"`

	// 最大长度。
	MaxLength *int32 `json:"max_length,omitempty"`
}

func (o AimTemplateParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimTemplateParams struct{}"
	}

	return strings.Join([]string{"AimTemplateParams", string(data)}, " ")
}
