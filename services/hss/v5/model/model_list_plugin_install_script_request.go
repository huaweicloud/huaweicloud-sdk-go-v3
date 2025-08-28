package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginInstallScriptRequest Request Object
type ListPluginInstallScriptRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 是否非华为云 **约束限制**： 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**： false
	OutsideHost *bool `json:"outside_host,omitempty"`

	// **参数解释**： 是否批量安装 **约束限制**： 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**： true
	BatchInstall *bool `json:"batch_install,omitempty"`

	// **参数解释**： 插件类型 **约束限制**： 不涉及 **取值范围**： - opa-docker-authz：docker插件。  **默认取值**： opa-docker-authz
	Plugin string `json:"plugin"`

	// **参数解释**： 操作类型 **约束限制**： 不涉及 **取值范围**： - install：安装。 - upgrade：升级。 - uninstall：卸载。  **默认取值**： install
	OperateType string `json:"operate_type"`
}

func (o ListPluginInstallScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginInstallScriptRequest struct{}"
	}

	return strings.Join([]string{"ListPluginInstallScriptRequest", string(data)}, " ")
}
