package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulContainerAppsRequest Request Object
type ListVulContainerAppsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 受影响容器id
	ContainerId string `json:"container_id"`

	// 漏洞ID
	VulId string `json:"vul_id"`

	// 漏洞的处理状态，包含如下：   -unhandled : 未处理   -handled : 已处理
	HandleStatus string `json:"handle_status"`

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListVulContainerAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulContainerAppsRequest struct{}"
	}

	return strings.Join([]string{"ListVulContainerAppsRequest", string(data)}, " ")
}
