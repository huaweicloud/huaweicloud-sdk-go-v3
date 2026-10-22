package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2EndpointRequestBody **参数解释**：更新端点请求。
type UpdateV2EndpointRequestBody struct {

	// **参数解释**：端点显示名称。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：资源描述信息。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：计算资源容量配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Capacity *[]V2EndpointCapacityItem `json:"capacity,omitempty"`

	Engine *V2EndpointUpdateEngineConfig `json:"engine,omitempty"`

	PrometheusConfig *PrometheusConfig `json:"prometheus_config,omitempty"`

	LtsConfig *LtsConfig `json:"lts_config,omitempty"`

	Credential *CredentialConfig `json:"credential,omitempty"`

	// **参数解释**：标签列表，每个元素为一个V2ResourceTag对象，包含key（标签键）和value（标签值）字段。 **约束限制**：最多20个标签。
	Tags *[]V2ResourceTag `json:"tags,omitempty"`
}

func (o UpdateV2EndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2EndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateV2EndpointRequestBody", string(data)}, " ")
}
