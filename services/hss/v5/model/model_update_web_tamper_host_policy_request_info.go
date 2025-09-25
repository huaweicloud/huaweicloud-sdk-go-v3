package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebTamperHostPolicyRequestInfo 网页防篡改策略信息
type UpdateWebTamperHostPolicyRequestInfo struct {
	ProtectDirInfo *WebTamperProtectDirRequestInfo `json:"protect_dir_info"`

	// **参数解释**: 定时开关设置状态 **约束限制**: 不涉及 **取值范围**: - True ：开启定时关闭防护功能，必须填写timing_off_config_info。 - False ：关闭定时关闭防护功能，无需填写timing_off_config_info。  **默认取值**: False
	EnableTimingOff *bool `json:"enable_timing_off,omitempty"`

	TimingOffConfigInfo *WebTamperTimingOffConfigInfoRequestInfo `json:"timing_off_config_info,omitempty"`

	// **参数解释**: 动态网页防篡改开启状态，仅Linux服务器支持。 **约束限制**: 仅Linux服务器支持开启动态网页防篡改，Windows服务器不可填写该字段。 **取值范围**: - True ：开启动态网页防篡改，必须填写rasp_path。 - False ：关闭动态网页防篡改，无需填写rasp_path。  **默认取值**: False
	EnableRaspProtect *bool `json:"enable_rasp_protect,omitempty"`

	// **参数解释**: 动态网页防篡改的Tomcat bin目录，仅Linux服务器支持。 **约束限制**: 仅Linux服务器支持配置动态网页防篡改的Tomcat bin目录，Windows服务器不可填写该字段。 **取值范围**: 字符长度1-256位，必须以/开头，不能以/结尾，只能包含英文大小写字母，数字，下划线，中划线和点。 **默认取值**: 不涉及
	RaspPath *string `json:"rasp_path,omitempty"`

	// **参数解释**: 特权进程开启状态 **约束限制**: 不涉及 **取值范围**: - True ：开启特权进程，必须填写privileged_process_info。 - False ：关闭特权进程，无需填写privileged_process_info。  **默认取值**: False
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
