package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VaultBatchUpdate 存储库批量修改属性
type VaultBatchUpdate struct {

	// 存储库smn消息通知开关
	SmnNotify *bool `json:"smn_notify,omitempty"`

	// 存储库容量阈值
	Threshold *int32 `json:"threshold,omitempty"`
}

func (o VaultBatchUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBatchUpdate struct{}"
	}

	return strings.Join([]string{"VaultBatchUpdate", string(data)}, " ")
}
