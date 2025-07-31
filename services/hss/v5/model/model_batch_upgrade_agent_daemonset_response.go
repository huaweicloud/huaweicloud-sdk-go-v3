package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeAgentDaemonsetResponse Response Object
type BatchUpgradeAgentDaemonsetResponse struct {

	// 批量升级列表
	DataList       *[]UpdateDaemonsetRespInfo `json:"data_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchUpgradeAgentDaemonsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeAgentDaemonsetResponse struct{}"
	}

	return strings.Join([]string{"BatchUpgradeAgentDaemonsetResponse", string(data)}, " ")
}
