package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FeatureSetOpenApiVo 特性集OpenApi返回体
type FeatureSetOpenApiVo struct {

	// 特性集ID
	Id *string `json:"id,omitempty"`

	// 编号
	Number *string `json:"number,omitempty"`

	// 父特性集ID
	ParentId *string `json:"parent_id,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 位置信息
	PositionFloat *float32 `json:"position_float,omitempty"`

	CreatedBy *UserEntity `json:"created_by,omitempty"`

	ModifiedBy *UserEntity `json:"modified_by,omitempty"`

	// **参数解释**： 特性集创建时间的时间戳。 **取值范围**： 不涉及。
	CreatedDate *string `json:"created_date,omitempty"`

	// **参数解释**： 特性集修改时间的时间戳。 **取值范围**： 不涉及。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 子特性集
	ChildFs *[]FeatureSetOpenApiVo `json:"child_fs,omitempty"`
}

func (o FeatureSetOpenApiVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureSetOpenApiVo struct{}"
	}

	return strings.Join([]string{"FeatureSetOpenApiVo", string(data)}, " ")
}
