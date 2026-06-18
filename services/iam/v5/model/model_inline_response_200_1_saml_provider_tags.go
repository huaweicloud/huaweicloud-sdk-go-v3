package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InlineResponse2001SamlProviderTags **参数解释**： 自定义标签信息。  **取值范围**： 不涉及。
type InlineResponse2001SamlProviderTags struct {

	// **参数解释**： 标签键。  **取值范围**： 字符串长度在 1 到 64 之间，可以包含任意语种字母、数字、空格以及\"_\"、\".\"、\":\"、\"=\"、\"+\"、\"-\"、\"@\"符号；首尾不能包含空格，且不能以\"_sys_\"开头。
	TagKey string `json:"tag_key"`

	// **参数解释**： 标签值。  **取值范围**： 字符串长度在 0 到 128 之间，可以包含任意语种字母、数字、空格以及\"_\"、\".\"、\":\"、\"/\"、\"=\"、\"+\"、\"-\"、\"@\"符号，可以为空字符串。
	TagValue string `json:"tag_value"`
}

func (o InlineResponse2001SamlProviderTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InlineResponse2001SamlProviderTags struct{}"
	}

	return strings.Join([]string{"InlineResponse2001SamlProviderTags", string(data)}, " ")
}
