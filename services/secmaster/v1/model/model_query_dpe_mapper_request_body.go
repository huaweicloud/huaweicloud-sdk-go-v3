package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDpeMapperRequestBody 查询分类映射请求体
type QueryDpeMapperRequestBody struct {

	// 映射id
	MappingId string `json:"mapping_id"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 映射id
	DatasourceInstanceId *string `json:"datasource_instance_id,omitempty"`

	// 状态(enabled，disabled)
	Status *string `json:"status,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 是否配置预处理规则
	HasPreprocessRule *bool `json:"has_preprocess_rule,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 0-10000 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 当前页码 **约束限制**: 不涉及
	Limit *int32 `json:"limit,omitempty"`
}

func (o QueryDpeMapperRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDpeMapperRequestBody struct{}"
	}

	return strings.Join([]string{"QueryDpeMapperRequestBody", string(data)}, " ")
}
