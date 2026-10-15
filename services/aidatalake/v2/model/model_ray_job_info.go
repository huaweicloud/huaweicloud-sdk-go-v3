package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayJobInfo **参数解释**：作业详情信息。 **约束限制**：不涉及。
type RayJobInfo struct {

	// **参数解释**：Job名称。 **约束限制**：不涉及。 **取值范围**：长度为1~47的英文字母、数字、中划线的组合。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：端点名称。 **约束限制**：不涉及。 **取值范围**：长度为1~63个字符。包含小写字母、数字、中划线的组合。字母开头、字母或数字结尾。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	Config *RayJobConfig `json:"config"`

	// **参数解释**：描述信息。 **约束限制**：不涉及。 **取值范围**：长度为0-1024的任意字符组合，不包含“!”、“<”、“>”、“=”、“&”、“\"”、“'”。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：作业ID。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：运行时长，单位：秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：Ray Job的唯一ID标识。 **约束限制**：不涉及。 **取值范围**：长度为0~64个字符。 **默认取值**：不涉及。
	SubmissionId *string `json:"submission_id,omitempty"`

	// **参数解释**：作业对应的端点类型。 **约束限制**：不涉及。 **取值范围**：   - RAY_JOB: RayJob类型的端点作业；   - RAY_CLUSTER: RayCluster类型的端点作业；   - RAY_FEDERATION: Ray联邦端点作业。 **默认取值**：不涉及。
	EndpointType *string `json:"endpoint_type,omitempty"`

	// **参数解释**：端点ID。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：联邦ID。 **约束限制**：如果作业未提交到联邦端点，该字段为空字符串。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	FederationId *string `json:"federation_id,omitempty"`

	// **参数解释**：联邦名称。 **约束限制**：不涉及。 **取值范围**：长度为1~63个字符。 **默认取值**：不涉及。
	FederationName *string `json:"federation_name,omitempty"`

	// **参数解释**：创建时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：更新时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// **参数解释**：作业开始执行的时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// **参数解释**：结束时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Status *StatusEnum `json:"status,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`

	Error *ErrorMessage `json:"error,omitempty"`

	// **参数解释**：指示该作业是否启用了历史服务器（History Server）功能。 **约束限制**：不涉及。 **取值范围**：   - true：启用History Server。   - false：禁用History Server。 **默认取值**：false。
	HistoryServerEnabled *bool `json:"history_server_enabled,omitempty"`
}

func (o RayJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayJobInfo struct{}"
	}

	return strings.Join([]string{"RayJobInfo", string(data)}, " ")
}
