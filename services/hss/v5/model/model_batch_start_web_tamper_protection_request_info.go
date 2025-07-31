package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartWebTamperProtectionRequestInfo 开启网页防篡改防护信息
type BatchStartWebTamperProtectionRequestInfo struct {

	// 主机ID数组，不能为空
	HostIdList []string `json:"host_id_list"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 计费模式   - packet_cycle: 包周期
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源标签列表
	Tags *[]TagInfo `json:"tags,omitempty"`

	ProtectDirInfo *WebTamperProtectDirRequestInfo `json:"protect_dir_info,omitempty"`

	// 定时开关状态
	EnableTimingOff *bool `json:"enable_timing_off,omitempty"`

	TimingOffConfigInfo *WebTamperTimingOffConfigInfoRequestInfo `json:"timing_off_config_info,omitempty"`

	// 动态网页防篡改开启状态
	EnableRaspProtect *bool `json:"enable_rasp_protect,omitempty"`

	// rasp path
	RaspPath *string `json:"rasp_path,omitempty"`

	// 特权进程状态
	EnablePrivilegedProcess *bool `json:"enable_privileged_process,omitempty"`

	PrivilegedProcessInfo *WebTamperPrivilegedProcessRequestInfo `json:"privileged_process_info,omitempty"`
}

func (o BatchStartWebTamperProtectionRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartWebTamperProtectionRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchStartWebTamperProtectionRequestInfo", string(data)}, " ")
}
