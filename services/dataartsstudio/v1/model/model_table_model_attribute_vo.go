package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableModelAttributeVo struct {

	// 编码。
	Id *string `json:"id,omitempty"`

	// 字段名。
	NameEn string `json:"name_en"`

	// 业务属性。
	NameCh string `json:"name_ch"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// obs路径，子路径。
	ObsLocation *string `json:"obs_location,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 字段类型。
	DataType string `json:"data_type"`

	DomainType *DataTypeDomainEnum `json:"domain_type,omitempty"`

	// 数据类型扩展字段。
	DataTypeExtend *string `json:"data_type_extend,omitempty"`

	// 是否主键。
	IsPrimaryKey bool `json:"is_primary_key"`

	// 是否分区键。
	IsPartitionKey *bool `json:"is_partition_key,omitempty"`

	// 是否外键。
	IsForeignKey *bool `json:"is_foreign_key,omitempty"`

	// 是否继承的属性。
	ExtendField *bool `json:"extend_field,omitempty"`

	// 是否不为空。
	NotNull *bool `json:"not_null,omitempty"`

	// 序号。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 所属关系建模的模型ID。
	TableModelId *int64 `json:"table_model_id,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 表标签。
	Tags *[]TagVo `json:"tags,omitempty"`

	// 关联的数据标准的ID。
	StandRowId *int64 `json:"stand_row_id,omitempty"`

	// 关联的数据标准名称。
	StandRowName *string `json:"stand_row_name,omitempty"`

	// 质量信息。
	QualityInfos *[]QualityInfoVo `json:"quality_infos,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o TableModelAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableModelAttributeVo struct{}"
	}

	return strings.Join([]string{"TableModelAttributeVo", string(data)}, " ")
}
