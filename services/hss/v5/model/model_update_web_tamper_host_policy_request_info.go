package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWebTamperHostPolicyRequestInfo struct {
	ProtectDirInfo *WebTamperProtectDirRequestInfo `json:"protect_dir_info,omitempty"`

	// 定时开关状态
	EnableTimingOff *bool `json:"enable_timing_off,omitempty"`

	TimingOffConfigInfo *WebTamperTimingOffConfigInfoRequestInfo `json:"timing_off_config_info,omitempty"`

	// 动态网页防篡改开启状态
	EnableRaspProtect *bool `json:"enable_rasp_protect,omitempty"`

	// 动态网页防篡改的Tomcat bin目录
	RaspPath *string `json:"rasp_path,omitempty"`

	// 特权进程状态
	EnablePrivilegedProcess *bool `json:"enable_privileged_process,omitempty"`

	PrivilegedProcessInfo *WebTamperPrivilegedProcessRequestInfo `json:"privileged_process_info,omitempty"`
}

func (o UpdateWebTamperHostPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperHostPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperHostPolicyRequestInfo", string(data)}, " ")
}
