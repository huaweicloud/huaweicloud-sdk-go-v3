package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkPolicyState 实例状态。
type CentralNetworkPolicyState struct {
	State *CentralNetworkPolicyStateEnum `json:"state"`
}

func (o CentralNetworkPolicyState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkPolicyState struct{}"
	}

	return strings.Join([]string{"CentralNetworkPolicyState", string(data)}, " ")
}
