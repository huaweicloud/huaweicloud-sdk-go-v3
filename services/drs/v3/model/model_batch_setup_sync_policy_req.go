package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetupSyncPolicyReq 批量设置同步策略请求体
type BatchSetupSyncPolicyReq struct {

	// **参数解释**：  批量设置同步策略请求列表，包含需要设置策略的同步任务信息。详情请参见SyncPolicyReq对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []SyncPolicyReq `json:"jobs"`
}

func (o BatchSetupSyncPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetupSyncPolicyReq struct{}"
	}

	return strings.Join([]string{"BatchSetupSyncPolicyReq", string(data)}, " ")
}
