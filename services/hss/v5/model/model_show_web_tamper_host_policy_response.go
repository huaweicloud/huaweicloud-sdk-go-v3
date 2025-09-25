package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebTamperHostPolicyResponse Response Object
type ShowWebTamperHostPolicyResponse struct {

	// **参数解释**: 防护目录数 **取值范围**: 最小值0，最大值50
	ProtectDirNum *int32 `json:"protect_dir_num,omitempty"`

	ProtectDirInfo *WtpProtectDirResponseInfo `json:"protect_dir_info,omitempty"`

	// **参数解释**: 定时开关状态 **取值范围**: - True ：开启定时关闭防护功能。 - False ：未开启定时关闭防护功能。
	EnableTimingOff *bool `json:"enable_timing_off,omitempty"`

	TimingOffConfigInfo *ListTimingOffConfigInfoResponseInfo `json:"timing_off_config_info,omitempty"`

	// **参数解释**: 动态网页防篡改开启状态，仅Linux服务器支持设置动态网页防篡改。 **取值范围**: - True ：开启动态网页防篡改防护。 - False ：未开启动态网页防篡改防护。
	EnableRaspProtect *bool `json:"enable_rasp_protect,omitempty"`

	// **参数解释**: 动态网页防篡改的Tomcat bin目录。 **取值范围**: 字符长度0-512位
	RaspPath *string `json:"rasp_path,omitempty"`

	// **参数解释**: 特权进程开启状态 **取值范围**: - True ：开启特权进程。 - False ：未开启特权进程。
	EnablePrivilegedProcess *bool `json:"enable_privileged_process,omitempty"`

	// **参数解释**: 特权进程子进程可信状态，需先开启特权进程 **取值范围**: - True ：开启特权进程子进程可信。 - False ：未开启特权进程子进程可信。
	PrivilegedChildStatus *bool `json:"privileged_child_status,omitempty"`

	// **参数解释**: 特权进程文件路径列表 **取值范围**: 最少0条，最多10条
	PrivilegedProcessPathList *[]string `json:"privileged_process_path_list,omitempty"`
	HttpStatusCode            int       `json:"-"`
}

func (o ShowWebTamperHostPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebTamperHostPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowWebTamperHostPolicyResponse", string(data)}, " ")
}
