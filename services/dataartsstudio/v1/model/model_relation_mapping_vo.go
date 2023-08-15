package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationMappingVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 关系id
	RelationId *int64 `json:"relation_id,omitempty"`

	// 源字段id
	SourceFieldId *int64 `json:"source_field_id,omitempty"`

	// 目标字段id
	TargetFieldId *int64 `json:"target_field_id,omitempty"`

	// 源表名称
	SourceFieldName *string `json:"source_field_name,omitempty"`

	// 目的表名称
	TargetFieldName *string `json:"target_field_name,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o RelationMappingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationMappingVo struct{}"
	}

	return strings.Join([]string{"RelationMappingVo", string(data)}, " ")
}
