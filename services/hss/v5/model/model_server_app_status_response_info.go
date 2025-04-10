package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerAppStatusResponseInfo struct {

	// web中间件名称
	WebMidware *string `json:"web_midware,omitempty"`

	// jdk版本
	JdkVersion *string `json:"jdk_version,omitempty"`

	// java应用监听的端口列表
	PortList *[]int32 `json:"port_list,omitempty"`

	// process id
	Pid *int32 `json:"pid,omitempty"`

	// 启动命令
	CmdLine *string `json:"cmd_line,omitempty"`

	// 动态接入报错信息
	ErrorInfo *string `json:"error_info,omitempty"`

	// 应用启动时间，毫秒级时间戳(ms)
	StartTime *string `json:"start_time,omitempty"`

	// java应用防护状态，包含如下4种。 - 0 ：未防护。 - 1 ：防护失败。 - 2 ：手动防护成功。 - 3 ：自动防护成功
	AppProtectStatus *int32 `json:"app_protect_status,omitempty"`

	// 检测规则标识
	ChkFeatureName *string `json:"chk_feature_name,omitempty"`
}

func (o ServerAppStatusResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAppStatusResponseInfo struct{}"
	}

	return strings.Join([]string{"ServerAppStatusResponseInfo", string(data)}, " ")
}
