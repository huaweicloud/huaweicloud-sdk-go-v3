package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebTamperHostPolicyResponse Response Object
type ShowWebTamperHostPolicyResponse struct {

	// 防护目录数
	ProtectDirNum *int32 `json:"protect_dir_num,omitempty"`

	ProtectDirInfo *WtpProtectDirResponseInfo `json:"protect_dir_info,omitempty"`

	// 定时开关状态
	EnableTimingOff *bool `json:"enable_timing_off,omitempty"`

	TimingOffConfigInfo *ListTimingOffConfigInfoResponseInfo `json:"timing_off_config_info,omitempty"`

	// 动态网页防篡改开启状态
	EnableRaspProtect *bool `json:"enable_rasp_protect,omitempty"`

	// rasp path
	RaspPath *string `json:"rasp_path,omitempty"`

	// 特权进程状态
	EnablePrivilegedProcess *bool `json:"enable_privileged_process,omitempty"`

	// 特权进程子进程可信状态
	PrivilegedChildStatus *bool `json:"privileged_child_status,omitempty"`

	// 特权进程路径集合
	PrivilegedProcessPathList *[]string `json:"privileged_process_path_list,omitempty"`

	PrivilegedProcessInfo *ListPrivilegedProcessResponseInfo `json:"privileged_process_info,omitempty"`
	HttpStatusCode        int                                `json:"-"`
}

func (o ShowWebTamperHostPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebTamperHostPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowWebTamperHostPolicyResponse", string(data)}, " ")
}
