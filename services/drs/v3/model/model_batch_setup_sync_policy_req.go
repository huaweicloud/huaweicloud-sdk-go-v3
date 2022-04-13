package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量设置同步策略请求体
type BatchSetupSyncPolicyReq struct {
	// 批量设置同步策略请求列表。

	Jobs []SyncPolicyReq `json:"jobs"`
}

func (o BatchSetupSyncPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetupSyncPolicyReq struct{}"
	}

	return strings.Join([]string{"BatchSetupSyncPolicyReq", string(data)}, " ")
}
