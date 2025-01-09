package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopLoginStatisticsRsp 登录状态统计响应。
type DesktopLoginStatisticsRsp struct {

	// 使用中。
	InUseNum int32 `json:"in_use_num"`

	// 关机数目(关机中、已关机)。
	StopNum int32 `json:"stop_num"`

	// 未注册数目。
	UnregisteredNum int32 `json:"unregistered_num"`

	// 未注册数目。
	UnableToConnectNum *int32 `json:"unable_to_connect_num,omitempty"`

	// 空闲数目。
	ReadyNum int32 `json:"ready_num"`

	// 断开连接数目。
	DisconnectedNum int32 `json:"disconnected_num"`
}

func (o DesktopLoginStatisticsRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopLoginStatisticsRsp struct{}"
	}

	return strings.Join([]string{"DesktopLoginStatisticsRsp", string(data)}, " ")
}
