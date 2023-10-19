package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkState 实例状态。
type CentralNetworkState struct {
	State *CentralNetworkStateEnum `json:"state"`
}

func (o CentralNetworkState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkState struct{}"
	}

	return strings.Join([]string{"CentralNetworkState", string(data)}, " ")
}
