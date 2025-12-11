package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationRequestBody struct {

	// **参数解释：** 参数模板名称。 **约束限制：** 最长64个字符，只允许大写字母、小写字母、数字、和“-_.”特殊字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 参数模板描述。 **约束限制：** 最长256个字符，不支持>!<\"&'=特殊字符。 **取值范围：** 不涉及。 **默认取值：** 空。
	Description *string `json:"description,omitempty"`

	Datastore *CreateConfigurationDatastoreOption `json:"datastore,omitempty"`

	// **参数解释：** 参数值对象，用户基于默认参数模板自定义的参数值。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 默认不修改参数值。
	Values map[string]string `json:"values,omitempty"`

	// **参数解释：** 实例ID。 **约束限制：** 实例ID可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 若传入此参数，则会基于此实例的参数信息创建参数模板。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o CreateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequestBody", string(data)}, " ")
}
