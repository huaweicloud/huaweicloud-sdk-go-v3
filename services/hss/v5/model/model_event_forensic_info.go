package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfo 事件取证
type EventForensicInfo struct {

	// **参数解释**： 发生时间，毫秒 **取值范围**： 不涉及
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 文件取证信息列表 **取值范围**： 不涉及
	FileForensicList *[]FileForensicInfo `json:"file_forensic_list,omitempty"`

	// **参数解释**： 进程链取证信息 **取值范围**： 不涉及
	ProcessChainForensic *[][]ProcessForensicInfo `json:"process_chain_forensic,omitempty"`

	NetworkForensic *EventForensicInfoNetworkForensic `json:"network_forensic,omitempty"`

	UserForensic *EventForensicInfoUserForensic `json:"user_forensic,omitempty"`

	RegistryForensic *EventForensicInfoRegistryForensic `json:"registry_forensic,omitempty"`

	AbnormalLoginForensic *EventForensicInfoAbnormalLoginForensic `json:"abnormal_login_forensic,omitempty"`

	ContainerForensic *EventForensicInfoContainerForensic `json:"container_forensic,omitempty"`

	MalwareForensic *EventForensicInfoMalwareForensic `json:"malware_forensic,omitempty"`

	AutoLaunchForensic *EventForensicInfoAutoLaunchForensic `json:"auto_launch_forensic,omitempty"`

	// **参数解释**： 内核取证信息 **取值范围**： 不涉及
	KernelForensicList *[]KernelForensicInfo `json:"kernel_forensic_list,omitempty"`

	GeoForensic *EventForensicInfoGeoForensic `json:"geo_forensic,omitempty"`

	StackForensic *EventForensicInfoStackForensic `json:"stack_forensic,omitempty"`

	ImageBlockForensic *EventForensicInfoImageBlockForensic `json:"image_block_forensic,omitempty"`

	// **参数解释**： 蜜罐相关取证信息 **取值范围**： 不涉及
	HoneyForensic *[]HoneyForensicInfo `json:"honey_forensic,omitempty"`
}

func (o EventForensicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfo struct{}"
	}

	return strings.Join([]string{"EventForensicInfo", string(data)}, " ")
}
