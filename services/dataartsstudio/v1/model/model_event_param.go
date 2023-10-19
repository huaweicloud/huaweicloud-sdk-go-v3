package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventParam struct {
	Guid *string `json:"guid,omitempty"`

	// 资产类型
	TypeName *string `json:"type_name,omitempty"`

	ModelId *string `json:"model_id,omitempty"`

	PropertyName *string `json:"property_name,omitempty"`

	PropertyValue *string `json:"property_value,omitempty"`

	// 搜索框输入
	Query *string `json:"query,omitempty"`

	Filter *DataMapFilterCriteria `json:"filter,omitempty"`

	GuidList *[]string `json:"guid_list,omitempty"`

	TraceId *string `json:"trace_id,omitempty"`

	SourceTraceId *string `json:"source_trace_id,omitempty"`

	MetadataTypeName *string `json:"metadata_type_name,omitempty"`

	SuperTypeNames *string `json:"super_type_names,omitempty"`

	WorkspaceIds *[]string `json:"workspace_ids,omitempty"`
}

func (o EventParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventParam struct{}"
	}

	return strings.Join([]string{"EventParam", string(data)}, " ")
}
