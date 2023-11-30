package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HierarchiesAttrVo 层级属性
type HierarchiesAttrVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 层级id
	HierarchiesId *int64 `json:"hierarchies_id,omitempty"`

	// 属性id
	AttrId *int64 `json:"attr_id,omitempty"`

	// 层次
	Level *int32 `json:"level,omitempty"`

	// 引用属性编码
	AttrNameEn *string `json:"attr_name_en,omitempty"`

	// 引用属性名称
	AttrNameCh *string `json:"attr_name_ch,omitempty"`

	// 详情属性id
	DetailAttrIds *[]int64 `json:"detail_attr_ids,omitempty"`

	// 详情属性英文
	DetailAttrNameEns *[]string `json:"detail_attr_name_ens,omitempty"`

	// 详情属性中文
	DetailAttrNameChs *[]string `json:"detail_attr_name_chs,omitempty"`

	Attr *DimensionAttributeVo `json:"attr,omitempty"`

	// 详情字段
	DetailAttrs *[]DimensionAttributeVo `json:"detail_attrs,omitempty"`
}

func (o HierarchiesAttrVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HierarchiesAttrVo struct{}"
	}

	return strings.Join([]string{"HierarchiesAttrVo", string(data)}, " ")
}
