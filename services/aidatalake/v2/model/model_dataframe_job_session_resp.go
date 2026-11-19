package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataframeJobSessionResp Aura Dataframe Job 详情 - session 列表响应
type DataframeJobSessionResp struct {

	// **参数解释**：Session ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：Session 状态。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：资源创建时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。 **取值范围**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：结束时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// **参数解释**：该 session 下关联的 statement 列表。 **约束限制**：session 无 statement 时返回空数组。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Statements *[]DataframeJobStatementResp `json:"statements,omitempty"`
}

func (o DataframeJobSessionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataframeJobSessionResp struct{}"
	}

	return strings.Join([]string{"DataframeJobSessionResp", string(data)}, " ")
}
