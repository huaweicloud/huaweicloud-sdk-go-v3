package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulnerabilitiesRequest Request Object
type ListVulnerabilitiesRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 漏洞类型，包含如下：   -linux_vul : linux漏洞   -windows_vul : windows漏洞   -web_cms : Web-CMS漏洞   -app_vul : 应用漏洞
	Type *string `json:"type,omitempty"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty"`

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 修复优先级 Critical 紧急 High  高 Medium 中 Low 低
	RepairPriority *string `json:"repair_priority,omitempty"`

	// 处置状态，包含如下:   - unhandled ：未处理   - handled : 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 漏洞编号
	CveId *string `json:"cve_id,omitempty"`

	// 漏洞标签
	LabelList *string `json:"label_list,omitempty"`

	// 漏洞状态
	Status *string `json:"status,omitempty"`

	// 资产重要性 important common test
	AssetValue *string `json:"asset_value,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`
}

func (o ListVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListVulnerabilitiesRequest", string(data)}, " ")
}
