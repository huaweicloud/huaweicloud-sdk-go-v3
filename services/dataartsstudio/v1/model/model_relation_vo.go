package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 源表ID，填写String类型替代Long类型。
	SourceTableId *string `json:"source_table_id,omitempty"`

	// 目标表ID，填写String类型替代Long类型。
	TargetTableId *string `json:"target_table_id,omitempty"`

	// 关系名称。
	Name string `json:"name"`

	// 源表名称。
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 目的表名称。
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 角色。
	Role *string `json:"role,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	SourceType *RelationType `json:"source_type,omitempty"`

	TargetType *RelationType `json:"target_type,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 表属性信息，只读。
	Mappings *[]RelationMappingVo `json:"mappings,omitempty"`
}

func (o RelationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationVo struct{}"
	}

	return strings.Join([]string{"RelationVo", string(data)}, " ")
}
