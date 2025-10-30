package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalAssetScanTaskRequestInfo **参数解释**: 下发任务的列表，三种参数选择其中一种即可，有多个参数则优先级顺序为all_hosts，server_group，host_ids **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
type CreateGlobalAssetScanTaskRequestInfo struct {

	// **参数解释**: 下发任务的主机列表 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值200 **默认取值**: 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**: 下发任务的主机组列表 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值200 **默认取值**: 不涉及
	ServerGroup *[]string `json:"server_group,omitempty"`

	// **参数解释**: 下发全部主机的扫描 **约束限制**: 不涉及 **取值范围**: - true：扫描全部主机 - false：不扫描全部主机 **默认取值**: 不涉及
	AllHosts *bool `json:"all_hosts,omitempty"`
}

func (o CreateGlobalAssetScanTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalAssetScanTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateGlobalAssetScanTaskRequestInfo", string(data)}, " ")
}
