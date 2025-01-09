package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MonitorUserOnlineInfo 桌面监控用户在线状态信息
type MonitorUserOnlineInfo struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`

	// 建立连接时间
	ConnectionSetupTime *string `json:"connection_setup_time,omitempty"`

	// 结束连接时间
	ConnectionEndTime *string `json:"connection_end_time,omitempty"`
}

func (o MonitorUserOnlineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorUserOnlineInfo struct{}"
	}

	return strings.Join([]string{"MonitorUserOnlineInfo", string(data)}, " ")
}
