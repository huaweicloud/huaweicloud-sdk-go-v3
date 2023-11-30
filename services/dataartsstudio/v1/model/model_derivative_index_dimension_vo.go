package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DerivativeIndexDimensionVo 维度分组(颗粒度)
type DerivativeIndexDimensionVo struct {

	// 维度分组id
	GroupId string `json:"group_id"`

	// 维度角色
	Role *string `json:"role,omitempty"`

	// 维度id
	DimensionId *int64 `json:"dimension_id,omitempty"`

	// 维度层级id
	HierarchiesId *int64 `json:"hierarchies_id,omitempty"`

	// 序号
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 维度分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 维度分组编码
	GroupCode *string `json:"group_code,omitempty"`

	BizType *BizTypeEnum `json:"biz_type"`

	// 层级属性
	Hierarchies *[]DimensionHierarchiesVo `json:"hierarchies,omitempty"`

	// 主题域分组中文名
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名
	L3 *string `json:"l3,omitempty"`

	// 主题域分组id
	L1Id *int64 `json:"l1_id,omitempty"`

	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象id
	L3Id *int64 `json:"l3_id,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// id
	Id *int64 `json:"id,omitempty"`
}

func (o DerivativeIndexDimensionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DerivativeIndexDimensionVo struct{}"
	}

	return strings.Join([]string{"DerivativeIndexDimensionVo", string(data)}, " ")
}
