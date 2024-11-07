package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkState 分支网络状态。
type SiteNetworkState struct {
	State *SiteNetworkStateEnum `json:"state"`
}

func (o SiteNetworkState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkState struct{}"
	}

	return strings.Join([]string{"SiteNetworkState", string(data)}, " ")
}
