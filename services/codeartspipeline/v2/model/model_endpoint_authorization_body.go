package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointAuthorizationBody **参数解释**： 鉴权信息。 **取值范围**： 不涉及。
type EndpointAuthorizationBody struct {

	// **参数解释**： 鉴权参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Parameters *interface{} `json:"parameters,omitempty"`

	// **参数解释**： 鉴权模式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Scheme *string `json:"scheme,omitempty"`
}

func (o EndpointAuthorizationBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointAuthorizationBody struct{}"
	}

	return strings.Join([]string{"EndpointAuthorizationBody", string(data)}, " ")
}
