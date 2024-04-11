package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationMappingVo struct {

	// 编码。
	Id *int64 `json:"id,omitempty"`

	// 关系ID。
	RelationId *int64 `json:"relation_id,omitempty"`

	// 源字段ID。
	SourceFieldId *int64 `json:"source_field_id,omitempty"`

	// 目标字段ID。
	TargetFieldId *int64 `json:"target_field_id,omitempty"`

	// 源表名称。
	SourceFieldName *string `json:"source_field_name,omitempty"`

	// 目的表名称。
	TargetFieldName *string `json:"target_field_name,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o RelationMappingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationMappingVo struct{}"
	}

	return strings.Join([]string{"RelationMappingVo", string(data)}, " ")
}
