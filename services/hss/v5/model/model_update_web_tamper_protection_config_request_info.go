package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWebTamperProtectionConfigRequestInfo struct {

	// **参数解释**: 网页防篡改的类型 **约束限制**: 不涉及 **取值范围**: - container_wtp：容器网页防篡改  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 防护配置id **约束限制**: 不涉及。 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ProtectionConfigId string `json:"protection_config_id"`

	// **参数解释**: 防护模式 **约束限制**: 不涉及 **取值范围**: - alarm：告警模式 - block：拦截模式  **默认取值**: alarm
	ProtectMode *string `json:"protect_mode,omitempty"`

	// **参数解释**: 是否对篡改行为进行进程监控 **约束限制**: 不涉及 **取值范围**: - true：监控篡改进程 - false：不监控篡改进程 **默认取值**: true
	MonitorProcess *bool `json:"monitor_process,omitempty"`

	// **参数解释**: 特权进程路径列表 **约束限制**: monitor_process值为true时生效（不进行进程监控无法判断特权进程） **取值范围**: 最少0条，最多10条 **默认取值**: 不涉及
	PrivilegedProcessList *[]string `json:"privileged_process_list,omitempty"`

	// **参数解释**: 特权进程的子进程是否仍享有特权 **约束限制**: monitor_process值为true时生效（不进行进程监控无法判断特权进程） **取值范围**: - true：特权进程的子进程依然是特权进程 - false：特权进程的子进程不再是特权进程 **默认取值**: false
	PrivilegedChild *bool `json:"privileged_child,omitempty"`

	ContainerWtpInfo *UpdateWebTamperProtectionConfigRequestInfoContainerWtpInfo `json:"container_wtp_info,omitempty"`

	// **参数解释**： 防护目录信息 **约束限制**: 不涉及 **取值范围**： 最少1条，最多50条 **默认取值**: 不涉及
	ProtectDirectoryInfoList *[]ProtectDirectoryInfo `json:"protect_directory_info_list,omitempty"`

	// **参数解释**： 防护排除的文件类型列表 **约束限制**: 不涉及 **取值范围**： 最少0条，最多10条 **默认取值**: 不涉及
	ExcludeFileTypes *[]string `json:"exclude_file_types,omitempty"`
}

func (o UpdateWebTamperProtectionConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperProtectionConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperProtectionConfigRequestInfo", string(data)}, " ")
}
