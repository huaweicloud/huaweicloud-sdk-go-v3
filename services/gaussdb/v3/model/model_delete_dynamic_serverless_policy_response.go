package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDynamicServerlessPolicyResponse Response Object
type DeleteDynamicServerlessPolicyResponse struct {

	// **参数解释**：  删除动态Serverless算力策略的任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDynamicServerlessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDynamicServerlessPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteDynamicServerlessPolicyResponse", string(data)}, " ")
}
