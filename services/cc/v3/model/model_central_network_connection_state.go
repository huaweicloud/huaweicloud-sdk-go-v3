package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkConnectionState 实例状态。
type CentralNetworkConnectionState struct {
	State *CentralNetworkConnectionStateEnum `json:"state"`
}

func (o CentralNetworkConnectionState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkConnectionState struct{}"
	}

	return strings.Join([]string{"CentralNetworkConnectionState", string(data)}, " ")
}
