package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncClusterProtectionEventResponse Response Object
type SyncClusterProtectionEventResponse struct {

	// 创建同步任务是否成功
	Synched        *bool `json:"synched,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SyncClusterProtectionEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncClusterProtectionEventResponse struct{}"
	}

	return strings.Join([]string{"SyncClusterProtectionEventResponse", string(data)}, " ")
}
