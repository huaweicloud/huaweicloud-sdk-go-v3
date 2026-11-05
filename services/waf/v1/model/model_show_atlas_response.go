package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAtlasResponse Response Object
type ShowAtlasResponse struct {

	// 图库ID
	Id *string `json:"id,omitempty"`

	// 图库名称
	Name *string `json:"name,omitempty"`

	// 图库类型 0=白库（豁免/可信样本） / 1=黑库（敏感/拦截样本）
	Type *int32 `json:"type,omitempty"`

	// 图库描述
	Description *string `json:"description,omitempty"`

	// 创建时间（毫秒）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间（毫秒）
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 图库下已加入（status≥0）且未删除的图片数量；创建返回 0
	ImageCount *int64 `json:"image_count,omitempty"`

	// 引用此图库的规则列表
	ReferencedRules *[]ReferencedRule `json:"referenced_rules,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ShowAtlasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAtlasResponse struct{}"
	}

	return strings.Join([]string{"ShowAtlasResponse", string(data)}, " ")
}
