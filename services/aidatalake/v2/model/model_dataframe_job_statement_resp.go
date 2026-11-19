package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataframeJobStatementResp Aura Dataframe Job 详情 - statement 列表响应
type DataframeJobStatementResp struct {

	// **参数解释**：Statement ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：执行状态。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：SQL 执行内容。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StatementContent *string `json:"statement_content,omitempty"`

	// **参数解释**：已处理条数。 **约束限制**：由 operator_realtime 上报更新。 **取值范围**：不涉及。 **默认取值**：0。
	TupleProcessed *int64 `json:"tuple_processed,omitempty"`

	// **参数解释**：资源创建时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。 **取值范围**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：结束时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FinishTime *sdktime.SdkTime `json:"finish_time,omitempty"`
}

func (o DataframeJobStatementResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataframeJobStatementResp struct{}"
	}

	return strings.Join([]string{"DataframeJobStatementResp", string(data)}, " ")
}
