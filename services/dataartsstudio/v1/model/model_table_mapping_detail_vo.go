package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableMappingDetailVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 名称。
	MappingId *string `json:"mapping_id,omitempty"`

	// 目的字段ID，填写String类型替代Long类型。
	TargetAttrId *string `json:"target_attr_id,omitempty"`

	// 目的字段排序。
	TargetAttrName string `json:"target_attr_name"`

	// 源表ID。
	SrcTableIds *string `json:"src_table_ids,omitempty"`

	// 源表名称数组，只读。
	SrcTableNames *[]string `json:"src_table_names,omitempty"`

	// 源表db名称数组，只读。
	SrcTableDbNames *[]string `json:"src_table_db_names,omitempty"`

	// 源表在关系建模中的模型ID数组，只读，填写String类型替代Long类型。
	SrcTableModelIds *[]string `json:"src_table_model_ids,omitempty"`

	// 源表ID数组，只读，填写String类型替代Long类型。
	SrcTableIdList *[]string `json:"src_table_id_list,omitempty"`

	// 源表字段ID。
	SrcAttrIds *string `json:"src_attr_ids,omitempty"`

	// 源表字段名称数组，只读。
	SrcAttrNames *[]string `json:"src_attr_names,omitempty"`

	// 源表字段ID数组，只读，填写String类型替代Long类型。
	SrcAttrIdList *[]string `json:"src_attr_id_list,omitempty"`

	// 备注。
	Remark *string `json:"remark,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 目标属性。
	TargetAttr *interface{} `json:"target_attr,omitempty"`
}

func (o TableMappingDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableMappingDetailVo struct{}"
	}

	return strings.Join([]string{"TableMappingDetailVo", string(data)}, " ")
}
