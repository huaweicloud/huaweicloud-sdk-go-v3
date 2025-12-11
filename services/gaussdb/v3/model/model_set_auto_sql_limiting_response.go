package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoSqlLimitingResponse Response Object
type SetAutoSqlLimitingResponse struct {

	// **参数解释**：  开启自治限流任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAutoSqlLimitingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoSqlLimitingResponse struct{}"
	}

	return strings.Join([]string{"SetAutoSqlLimitingResponse", string(data)}, " ")
}
