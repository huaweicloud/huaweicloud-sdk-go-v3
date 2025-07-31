package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFilesRequest Request Object
type ListIacFilesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 扫描方式 **约束限制**: 不涉及 **取值范围**: - manual_scan：手动扫描 - cicd_scan：cicd扫描 **默认取值**: 不涉及
	ScanType string `json:"scan_type"`

	// **参数解释**: 文件ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64  **默认取值**: 不涉及
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**: 文件名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-512  **默认取值**: 不涉及
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 文件类型 **约束限制**: 不涉及 **取值范围**: - dockerfile：Dockerfile文件 - k8s_yaml：k8s yaml文件  **默认取值**: 不涉及
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**: 是否有风险 **约束限制**: 不涉及 **取值范围**: - true：有风险 - false：无风险  **默认取值**: 不涉及
	Risky *bool `json:"risky,omitempty"`

	// **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	CicdId *string `json:"cicd_id,omitempty"`

	// **参数解释**: CI/CD名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	CicdName *string `json:"cicd_name,omitempty"`
}

func (o ListIacFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFilesRequest struct{}"
	}

	return strings.Join([]string{"ListIacFilesRequest", string(data)}, " ")
}
