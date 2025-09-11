package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddHbaConfsResponse Response Object
type AddHbaConfsResponse struct {

	// **参数解释**: 结果码。 **取值范围**: 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**: 结果描述。 **取值范围**: 不涉及。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddHbaConfsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHbaConfsResponse struct{}"
	}

	return strings.Join([]string{"AddHbaConfsResponse", string(data)}, " ")
}
