package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebTamperProtectionConfigResponseInfo struct {

	// **参数解释**: 网页防篡改配置id **取值范围**: 字符长度1-64位
	Id *string `json:"id,omitempty"`

	// **参数解释**: 防护模式 **取值范围**: - alarm：告警模式 - block：拦截模式
	ProtectMode *string `json:"protect_mode,omitempty"`

	// **参数解释**: 是否对篡改行为进行进程监控 **取值范围**: - true：监控篡改进程 - false：不监控篡改进程
	MonitorProcess *bool `json:"monitor_process,omitempty"`

	// **参数解释**: 特权进程路径列表 **取值范围**: 最少0条，最多10条
	PrivilegedProcessList *[]string `json:"privileged_process_list,omitempty"`

	// **参数解释**: 特权进程的子进程是否仍享有特权 **取值范围**: - true：特权进程的子进程依然是特权进程 - false：特权进程的子进程不再是特权进程
	PrivilegedChild *bool `json:"privileged_child,omitempty"`

	// **参数解释**： 防护目录信息 **取值范围**： 最少1条，最多50条
	ProtectDirectoryInfoList *[]ProtectDirectoryResponseInfo `json:"protect_directory_info_list,omitempty"`

	// **参数解释**： 防护排除的文件类型列表 **取值范围**： 最少0条，最多10条
	ExcludeFileTypes *[]string `json:"exclude_file_types,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: - not_protect：未防护 - protected：防护中 - partial_protected：部分防护 - protect_failed：防护失败 - redundant：防护冗余
	Status *string `json:"status,omitempty"`

	// **参数解释**: 防护中的容器数量 **取值范围**: 最小值0，最大值2147483647
	ProtectedContainerNums *int32 `json:"protected_container_nums,omitempty"`

	// **参数解释**: 当前防护配置产生的防护事件数量 **取值范围**: 最小值0，最大值2147483647
	ProtectedEventNums *int32 `json:"protected_event_nums,omitempty"`

	// **参数解释**: 防护配置绑定的配额id **取值范围**: 字符长度1-64位
	ResourceId *string `json:"resource_id,omitempty"`

	ContainerWtpInfo *WebTamperProtectionConfigResponseInfoContainerWtpInfo `json:"container_wtp_info,omitempty"`
}

func (o WebTamperProtectionConfigResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionConfigResponseInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionConfigResponseInfo", string(data)}, " ")
}
