package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBlacklistEnableDto 包含打开或者关闭流量过滤功能的开关信息：打开或者关闭，后端 根据此参数来打开或者关闭功能。
type IpBlacklistEnableDto struct {

	// 打开或关闭IP黑名单功能 0：disable 1：enable
	Status int32 `json:"status"`
}

func (o IpBlacklistEnableDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBlacklistEnableDto struct{}"
	}

	return strings.Join([]string{"IpBlacklistEnableDto", string(data)}, " ")
}
