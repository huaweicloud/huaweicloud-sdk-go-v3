package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHbaConfResponse Response Object
type ModifyHbaConfResponse struct {

	// **参数解释**: 结果码。 **取值范围**: 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**: 结果描述。 **取值范围**: 不涉及。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyHbaConfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHbaConfResponse struct{}"
	}

	return strings.Join([]string{"ModifyHbaConfResponse", string(data)}, " ")
}
