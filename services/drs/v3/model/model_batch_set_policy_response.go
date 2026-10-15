package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetPolicyResponse Response Object
type BatchSetPolicyResponse struct {

	// 批量设置同步策略返回列表
	Results *[]SyncPolicyResp `json:"results,omitempty"`

	// **参数解释**：  返回的任务信息总数，表示本次查询成功返回的任务数量，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchSetPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetPolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchSetPolicyResponse", string(data)}, " ")
}
