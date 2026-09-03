package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiskSpaceDiagnosisRequestBody 磁盘容量趋势诊断请求体
type CreateDiskSpaceDiagnosisRequestBody struct {

	// **参数解释**：  开始日期，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。  **约束限制**：  开始日期范围需要在24小时内。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**：  结束日期，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。  **约束限制**：  开始日期范围需要在24小时内。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime string `json:"end_time"`
}

func (o CreateDiskSpaceDiagnosisRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiskSpaceDiagnosisRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDiskSpaceDiagnosisRequestBody", string(data)}, " ")
}
