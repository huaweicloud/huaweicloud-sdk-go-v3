package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogAnalysisResponse Response Object
type ListSlowLogAnalysisResponse struct {

	// **参数解释**：  分页参数: 每页记录数。  **参数范围**：  不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：  慢日志列表。  **参数范围**：  不涉及。
	SlowLogList    *[]EsdbSlowSqlTemplateItem `json:"slow_log_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSlowLogAnalysisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogAnalysisResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogAnalysisResponse", string(data)}, " ")
}
