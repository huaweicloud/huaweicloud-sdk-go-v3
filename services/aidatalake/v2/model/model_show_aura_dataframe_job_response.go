package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraDataframeJobResponse Response Object
type ShowAuraDataframeJobResponse struct {

	// **参数解释**：作业ID。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：作业名称。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：描述信息。 **约束限制**：不涉及。 **取值范围**：长度为0-1024的任意字符组合，不包含“!”、“<”、“>”、“=”、“&”、“\"”、“'”。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：作业状态。 **约束限制**：不涉及。 **取值范围**：长度为1~128的任意字符组合。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：作业启动命令。 **约束限制**：不涉及。 **取值范围**：长度为1~1024的任意字符组合。 **默认取值**：不涉及。
	Entrypoint *string `json:"entrypoint,omitempty"`

	// **参数解释**：运行时长，单位：秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Duration *int64 `json:"duration,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`

	// **参数解释**：创建时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：作业开始执行的时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// **参数解释**：结束时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	ExecutionResult *ExecutionResult `json:"execution_result,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o ShowAuraDataframeJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraDataframeJobResponse struct{}"
	}

	return strings.Join([]string{"ShowAuraDataframeJobResponse", string(data)}, " ")
}
