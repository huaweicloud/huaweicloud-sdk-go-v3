package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeUpgradeStrategy 节点更新策略
type NodeUpgradeStrategy struct {

	// 策略类型
	Type *string `json:"type,omitempty"`

	RollingUpdate *RollingUpdateNodeUpgradeStrategy `json:"rollingUpdate,omitempty"`
}

func (o NodeUpgradeStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeUpgradeStrategy struct{}"
	}

	return strings.Join([]string{"NodeUpgradeStrategy", string(data)}, " ")
}
