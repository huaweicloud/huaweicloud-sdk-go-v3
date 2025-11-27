package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperProtectionDirectoryRequest Request Object
type ListWebTamperProtectionDirectoryRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 网页防篡改类型 **约束限制**: 不涉及。 **取值范围**: - container_wtp：容器网页防篡改 **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 防护配置id **约束限制**: 不涉及。 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ProtectionConfigId string `json:"protection_config_id"`

	// **参数解释**: 受影响容器id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 防护状态 **约束限制**: type为container_wtp时必传 **取值范围**: - protected：防护中 - protect_failed：防护失败 **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 防护目录 **约束限制**: 不涉及。 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	ProtectDirectory *string `json:"protect_directory,omitempty"`
}

func (o ListWebTamperProtectionDirectoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperProtectionDirectoryRequest struct{}"
	}

	return strings.Join([]string{"ListWebTamperProtectionDirectoryRequest", string(data)}, " ")
}
