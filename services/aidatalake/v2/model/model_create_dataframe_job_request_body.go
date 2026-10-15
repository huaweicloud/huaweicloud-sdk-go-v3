package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDataframeJobRequestBody struct {

	// **参数解释**：作业名称。  **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：描述信息。 **约束限制**：不涉及。 **取值范围**：长度为0-1024的任意字符组合，不包含“!”、“<”、“>”、“=”、“&”、“\"”、“'”。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：Endpoint名称。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	// **参数解释**：委托名称。  **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	AgencyName *string `json:"agency_name,omitempty"`

	Config *DataframeJobConfig `json:"config,omitempty"`
}

func (o CreateDataframeJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataframeJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataframeJobRequestBody", string(data)}, " ")
}
