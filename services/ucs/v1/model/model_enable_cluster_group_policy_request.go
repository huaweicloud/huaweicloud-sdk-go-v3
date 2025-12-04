package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableClusterGroupPolicyRequest Request Object
type EnableClusterGroupPolicyRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	// 重试启动策略管理功能
	Retry *string `json:"retry,omitempty"`
}

func (o EnableClusterGroupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableClusterGroupPolicyRequest struct{}"
	}

	return strings.Join([]string{"EnableClusterGroupPolicyRequest", string(data)}, " ")
}
