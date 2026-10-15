package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRaySessionRequestBody **参数解释**：创建交互式会话的输入请求。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type CreateRaySessionRequestBody struct {

	// **参数解释**：Job名称。 **约束限制**：不涉及。 **取值范围**：长度为1~47的英文字母、数字、中划线的组合。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：描述信息。 **约束限制**：不涉及。 **取值范围**：长度为0-1024的任意字符组合，不包含“!”、“<”、“>”、“=”、“&”、“\"”、“'”。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：端点名称。 **约束限制**：遵循RFC 1035规范，字母开头，字母或数字结尾。 **取值范围**：长度为1~63个字符，包含小写字母、数字、中划线的组合。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`
}

func (o CreateRaySessionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRaySessionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRaySessionRequestBody", string(data)}, " ")
}
