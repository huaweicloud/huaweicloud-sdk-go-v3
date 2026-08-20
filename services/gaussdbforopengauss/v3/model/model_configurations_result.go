package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationsResult **参数解释**: 参数模板信息。
type ConfigurationsResult struct {

	// **参数解释**: 参数模板ID。参数模板的唯一标识。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。
	Id string `json:"id"`

	// **参数解释**: 参数模板名称。 **取值范围**: 参数模板名称在1到64个字符之间，区分大小写，可包含字母、数字、英文中划线、下划线或句点，不能包含其他特殊字符。
	Name string `json:"name"`

	// **参数解释**: 参数模板描述。 **取值范围**: 描述不能超过256个字符，且不能包含回车和 ! < \" = ' > &这些特殊字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**: 引擎版本。 **取值范围**: 不涉及。
	DatastoreVersion string `json:"datastore_version"`

	// **参数解释**: 引擎名称。 **取值范围**: GaussDB。
	DatastoreName string `json:"datastore_name"`

	// **参数解释**: 节点类型。 **取值范围**: - independent：独立部署。 - ha：集中式。 - combined：混合部署。
	NodeType string `json:"node_type"`

	// **参数解释**: 实例类型。 **取值范围**: - Enterprise：分布式实例（企业版）。 - centralization_standard：集中式版实例。  区分大小写。
	HaMode string `json:"ha_mode"`

	// **参数解释**: 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。例如：2024-07-03T14:18:55。 **取值范围**: 不涉及。
	Created string `json:"created"`

	// **参数解释**: 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。例如：2024-07-03T14:18:55。 **取值范围**: 不涉及。
	Updated string `json:"updated"`

	// **参数解释**: 是否是用户自定义参数模板。 **取值范围**: - false：表示为系统默认参数模板。 - true：表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined"`
}

func (o ConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ConfigurationsResult", string(data)}, " ")
}
