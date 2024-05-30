package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationMappingVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 关系ID，填写String类型替代Long类型。
	RelationId *string `json:"relation_id,omitempty"`

	// 源字段ID，填写String类型替代Long类型。
	SourceFieldId *string `json:"source_field_id,omitempty"`

	// 目标字段ID，填写String类型替代Long类型。
	TargetFieldId *string `json:"target_field_id,omitempty"`

	// 源表名称。
	SourceFieldName *string `json:"source_field_name,omitempty"`

	// 目的表名称。
	TargetFieldName *string `json:"target_field_name,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o RelationMappingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationMappingVo struct{}"
	}

	return strings.Join([]string{"RelationMappingVo", string(data)}, " ")
}
