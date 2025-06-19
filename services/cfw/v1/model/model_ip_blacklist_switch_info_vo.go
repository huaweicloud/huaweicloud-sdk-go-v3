package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBlacklistSwitchInfoVo 包含流量过滤功能的开关信息
type IpBlacklistSwitchInfoVo struct {

	// IP黑名单功能开关信息 1:enable 0:disable
	Status *int32 `json:"status,omitempty"`
}

func (o IpBlacklistSwitchInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBlacklistSwitchInfoVo struct{}"
	}

	return strings.Join([]string{"IpBlacklistSwitchInfoVo", string(data)}, " ")
}
