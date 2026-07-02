package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDynamicServerlessPolicyResponse Response Object
type UpdateDynamicServerlessPolicyResponse struct {

	// **参数解释**：  配置动态Serverless功能的任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDynamicServerlessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDynamicServerlessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDynamicServerlessPolicyResponse", string(data)}, " ")
}
