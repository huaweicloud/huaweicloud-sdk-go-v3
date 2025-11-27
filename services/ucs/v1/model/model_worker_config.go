package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkerConfig struct {

	// 节点个数
	Replicas *int32 `json:"replicas,omitempty"`

	Strategy *NodeUpgradeStrategy `json:"strategy,omitempty"`
}

func (o WorkerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkerConfig struct{}"
	}

	return strings.Join([]string{"WorkerConfig", string(data)}, " ")
}
