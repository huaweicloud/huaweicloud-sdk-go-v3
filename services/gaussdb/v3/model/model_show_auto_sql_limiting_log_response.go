package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSqlLimitingLogResponse Response Object
type ShowAutoSqlLimitingLogResponse struct {

	// **参数解释**：  查询自治限流执行记录列表。  **约束限制**：  不涉及。
	Logs           *[]AutoSqlLimitingLog `json:"logs,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowAutoSqlLimitingLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSqlLimitingLogResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSqlLimitingLogResponse", string(data)}, " ")
}
