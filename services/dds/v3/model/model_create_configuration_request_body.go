package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationRequestBody struct {

	// **参数解释：** 参数模板名称。 **约束限制：** 长度1到64位之间，区分大小写字母，可包含字母、数字、中划线、下划线或句点，不能包含其他特殊字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 参数模板描述。 **约束限制：** 长度不超过256位，且不能包含回车和>!<\"&'=特殊字符。 **取值范围：** 不涉及。 **默认取值：** 默认为空。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。 **约束限制：** 当未传入entity_id参数时，此参数必选。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ParameterValues map[string]string `json:"parameter_values,omitempty"`

	Datastore *DatastoreResult `json:"datastore,omitempty"`

	// **参数解释：** 实例ID或组ID或节点ID。可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 若传入此参数，则会基于此实例、组或节点的参数信息创建参数模板，将会忽略parameter_values和datastore传参。 **约束限制：** 不涉及。 **取值范围：** 当实例类型是集群，取值为shard组或config组的组ID、mongos节点的节点ID、只读节点的节点ID。 当实例类型是副本集，传值为实例ID或只读节点的节点ID。 当实例类型是单节点，传值为实例ID。 **默认取值：** 不涉及。
	EntityId *string `json:"entity_id,omitempty"`
}

func (o CreateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequestBody", string(data)}, " ")
}
