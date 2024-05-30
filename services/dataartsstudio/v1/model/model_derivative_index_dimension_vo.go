package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DerivativeIndexDimensionVo 维度分组(颗粒度)。
type DerivativeIndexDimensionVo struct {

	// 维度分组ID。
	GroupId string `json:"group_id"`

	// 维度角色。
	Role *string `json:"role,omitempty"`

	// 维度ID，填写String类型替代Long类型。
	DimensionId *string `json:"dimension_id,omitempty"`

	// 维度层级ID，填写String类型替代Long类型。
	HierarchiesId *string `json:"hierarchies_id,omitempty"`

	// 序号，只读。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 维度分组名称。
	GroupName *string `json:"group_name,omitempty"`

	// 维度分组编码。
	GroupCode *string `json:"group_code,omitempty"`

	BizType *BizTypeEnum `json:"biz_type"`

	// 层级属性，只读。
	Hierarchies *[]DimensionHierarchiesVo `json:"hierarchies,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 主题域分组ID，只读，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象ID，只读，填写String类型替代Long类型。
	L3Id *string `json:"l3_id,omitempty"`

	// 数据连接类型。
	DwType *string `json:"dw_type,omitempty"`

	// 层级的ID，只读，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`
}

func (o DerivativeIndexDimensionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DerivativeIndexDimensionVo struct{}"
	}

	return strings.Join([]string{"DerivativeIndexDimensionVo", string(data)}, " ")
}
