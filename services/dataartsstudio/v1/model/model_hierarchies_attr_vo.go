package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HierarchiesAttrVo 层级属性。
type HierarchiesAttrVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 层级ID，填写String类型替代Long类型。
	HierarchiesId *string `json:"hierarchies_id,omitempty"`

	// 属性ID，填写String类型替代Long类型。
	AttrId *string `json:"attr_id,omitempty"`

	// 层次。
	Level *int32 `json:"level,omitempty"`

	// 引用属性编码。
	AttrNameEn *string `json:"attr_name_en,omitempty"`

	// 引用属性名称，只读。
	AttrNameCh *string `json:"attr_name_ch,omitempty"`

	// 详情属性ID，填写String类型替代Long类型。
	DetailAttrIds *[]string `json:"detail_attr_ids,omitempty"`

	// 详情属性英文。
	DetailAttrNameEns *[]string `json:"detail_attr_name_ens,omitempty"`

	// 详情属性中文，只读。
	DetailAttrNameChs *[]string `json:"detail_attr_name_chs,omitempty"`

	Attr *DimensionAttributeVo `json:"attr,omitempty"`

	// 详情字段，只读。
	DetailAttrs *[]DimensionAttributeVo `json:"detail_attrs,omitempty"`
}

func (o HierarchiesAttrVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HierarchiesAttrVo struct{}"
	}

	return strings.Join([]string{"HierarchiesAttrVo", string(data)}, " ")
}
