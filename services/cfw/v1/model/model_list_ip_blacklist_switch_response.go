package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpBlacklistSwitchResponse Response Object
type ListIpBlacklistSwitchResponse struct {
	Data           *IpBlacklistSwitchInfoVo `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListIpBlacklistSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpBlacklistSwitchResponse struct{}"
	}

	return strings.Join([]string{"ListIpBlacklistSwitchResponse", string(data)}, " ")
}
