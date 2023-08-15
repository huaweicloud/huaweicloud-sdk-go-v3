package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableMappingDetailVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 名称
	MappingId *string `json:"mapping_id,omitempty"`

	// 目的字段id
	TargetAttrId *int64 `json:"target_attr_id,omitempty"`

	// 目的字段排序
	TargetAttrName string `json:"target_attr_name"`

	// 源表id, 通过,join
	SrcTableIds *string `json:"src_table_ids,omitempty"`

	// 源表名称数组
	SrcTableNames *[]string `json:"src_table_names,omitempty"`

	// 源表db名称数组
	SrcTableDbNames *[]string `json:"src_table_db_names,omitempty"`

	// 源表模型id数组
	SrcTableModelIds *[]int64 `json:"src_table_model_ids,omitempty"`

	// 源表id数组
	SrcTableIdList *[]int64 `json:"src_table_id_list,omitempty"`

	// 源表字段id, 通过,join
	SrcAttrIds *string `json:"src_attr_ids,omitempty"`

	// 源表字段名称数组
	SrcAttrNames *[]string `json:"src_attr_names,omitempty"`

	// 源表字段id数组
	SrcAttrIdList *[]int64 `json:"src_attr_id_list,omitempty"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 目标属性
	TargetAttr *interface{} `json:"target_attr,omitempty"`
}

func (o TableMappingDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableMappingDetailVo struct{}"
	}

	return strings.Join([]string{"TableMappingDetailVo", string(data)}, " ")
}
