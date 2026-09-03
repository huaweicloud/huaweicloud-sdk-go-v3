package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchMySqlProxySlowLogResponse Response Object
type SwitchMySqlProxySlowLogResponse struct {

	// **参数解释**：  下发请求的链路ID。该返回值不支持在任务中心查询相关任务，请使用状态码判断是否请求成功。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchMySqlProxySlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchMySqlProxySlowLogResponse struct{}"
	}

	return strings.Join([]string{"SwitchMySqlProxySlowLogResponse", string(data)}, " ")
}
