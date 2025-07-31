package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulContainersRequest Request Object
type ListVulContainersRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 漏洞ID
	VulId string `json:"vul_id"`

	// 漏洞类型   - linux_vul : 漏洞类型-linux漏洞   - windows_vul : 漏洞类型-windows漏洞   - web_cms : Web-CMS漏洞   - app_vul : 应用漏洞   - urgent_vul : 应急漏洞
	Type string `json:"type"`

	// 受影响容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 漏洞状态   - vul_status_unfix : 未处理   - vul_status_ignored : 已忽略   - vul_status_verified : 验证中   - vul_status_fixing : 修复中   - vul_status_fixed : 修复成功   - vul_status_reboot : 修复成功待重启   - vul_status_failed : 修复失败   - vul_status_fix_after_reboot : 请重启主机再次修复
	Status *string `json:"status,omitempty"`

	// 处置状态，包含如下:   - unhandled ：未处理   - handled : 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 危险程度 ，Critical，High，Medium，Low
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 首次扫描时间最小值
	MinScanTime *int64 `json:"min_scan_time,omitempty"`

	// 首次扫描时间最大值
	MaxScanTime *int64 `json:"max_scan_time,omitempty"`

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListVulContainersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulContainersRequest struct{}"
	}

	return strings.Join([]string{"ListVulContainersRequest", string(data)}, " ")
}
