package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchHostsProtectStatusRequest Request Object
type SwitchHostsProtectStatusRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SwitchHostsProtectStatusRequestInfo `json:"body,omitempty"`
}

func (o SwitchHostsProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHostsProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"SwitchHostsProtectStatusRequest", string(data)}, " ")
}
