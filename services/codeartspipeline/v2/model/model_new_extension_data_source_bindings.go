package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewExtensionDataSourceBindings struct {

	// **参数解释**： 目标地址。 **取值范围**： 不涉及。
	Target *string `json:"target,omitempty"`

	// **参数解释**： 扩展点id。 **取值范围**： 不涉及。
	EndpointId *string `json:"endpointId,omitempty"`

	// **参数解释**： 数据源绑定名称。 **取值范围**： 不涉及。
	DataSourceName *string `json:"data_source_name,omitempty"`
}

func (o NewExtensionDataSourceBindings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionDataSourceBindings struct{}"
	}

	return strings.Join([]string{"NewExtensionDataSourceBindings", string(data)}, " ")
}
