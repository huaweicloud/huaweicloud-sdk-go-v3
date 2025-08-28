package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IacRequestInfo IaC风险筛选条件
type IacRequestInfo struct {

	// **参数解释**: IaC文件ID列表 **约束限制**: 不涉及 **取值范围**: 1-200个items **默认取值**: 不涉及
	FileIdList *[]string `json:"file_id_list,omitempty"`

	// **参数解释**: 文件名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-64 **默认取值**: 不涉及
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 文件类型 **约束限制**: 不涉及 **取值范围**: - dockerfile：Dockerfile文件。 - k8s_yaml：k8s yaml文件。  **默认取值**: 不涉及
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**: 是否有风险 **约束限制**: 不涉及 **取值范围**: - true：有风险。 - false：无风险。  **默认取值**: 不涉及
	Risky *bool `json:"risky,omitempty"`

	// **参数解释**: 扫描方式 **约束限制**: 不涉及 **取值范围**: - manual_scan：手动扫描。 - cicd_scan：cicd扫描。  **默认取值**: manual_scan
	ScanType *string `json:"scan_type,omitempty"`
}

func (o IacRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IacRequestInfo struct{}"
	}

	return strings.Join([]string{"IacRequestInfo", string(data)}, " ")
}
