package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationTreeRequest Request Object
type ListOrganizationTreeRequest struct {

	// **参数解释**: 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。 **约束限制**: 不涉及 **取值范围**: 字符长度1-2048位 **默认取值**: 不涉及
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 控制是否忽略本地缓存，强制从组织服务同步最新的组织树信息； **约束限制**: 无特殊约束，按需选择是否强制同步； **取值范围**: true（强制同步）、false（使用本地缓存，默认） **默认取值**: false
	IsRefresh *bool `json:"is_refresh,omitempty"`
}

func (o ListOrganizationTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationTreeRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationTreeRequest", string(data)}, " ")
}
