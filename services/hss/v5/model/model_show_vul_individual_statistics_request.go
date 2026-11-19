package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulIndividualStatisticsRequest Request Object
type ShowVulIndividualStatisticsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 需要查询的漏洞统计数据类型列表，多个类型以逗号分隔。 **约束限制**: 必填 **取值范围**: - need_urgent_repair_num：需要紧急修复的漏洞数量 - unhandled_num：未完成修复的漏洞数量 - existed_vul_host_num：存在漏洞的服务器数量 - today_handled_num：今日处理的漏洞数量 - all_handled_num：累计处理的漏洞数量 - vul_library_info：漏洞库统计数据 - linux_vul_unhandled_num：未处置的Linux漏洞数量 - windows_vul_unhandled_num：未处置的Windows漏洞数量 - web_cms_vul_unhandled_num：未处置的Web CMS漏洞数量 - app_vul_unhandled_num：未处置的应用漏洞数量 - urgent_vul_unhandled_num：未处置的应急漏洞数量 - max_vul_handle_num：单次最大可处置的漏洞数量  **默认取值**: 不涉及
	StatisticsTypes []string `json:"statistics_types"`
}

func (o ShowVulIndividualStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulIndividualStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowVulIndividualStatisticsRequest", string(data)}, " ")
}
