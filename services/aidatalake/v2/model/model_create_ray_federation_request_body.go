package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRayFederationRequestBody **参数解释**：创建Ray联邦端点请求。 **约束限制**：name为必填字段。
type CreateRayFederationRequestBody struct {

	// **参数解释**：Ray联邦端点名称。 **约束限制**：不涉及。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：Ray联邦端点显示名称。 **约束限制**：不涉及。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。 **默认取值**：不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：Ray联邦端点描述。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~1024个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：RayCluster端点信息列表。 **约束限制**：每个RayCluster只能属于一个Ray联邦端点。 **取值范围**：0~50个RayCluster端点。 **默认取值**：不涉及。
	RayClusterList *[]RayFederationRayClusterRequest `json:"ray_cluster_list,omitempty"`
}

func (o CreateRayFederationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRayFederationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRayFederationRequestBody", string(data)}, " ")
}
