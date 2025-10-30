package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RescanVulScanTaskRequestInfo struct {

	// **参数解释**: 重新扫描的主机范围 **约束限制**: 不涉及 **取值范围**: - all_failed_host：当前任务下所有扫描失败的主机 - specific_host：指定主机  **默认取值**: 不涉及
	RangeType string `json:"range_type"`

	// **参数解释**: 主机ID列表 **约束限制**: range_type值为“specific_host”时有效 **取值范围**: 最少1条，最多1000条  **默认取值**: 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`
}

func (o RescanVulScanTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RescanVulScanTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"RescanVulScanTaskRequestInfo", string(data)}, " ")
}
