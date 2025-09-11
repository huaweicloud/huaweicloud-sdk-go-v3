package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreHbaInfoResponse Response Object
type RestoreHbaInfoResponse struct {

	// **参数解释**: 结果码。 **取值范围**: 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**: 结果描述。 **取值范围**: 不涉及。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreHbaInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreHbaInfoResponse struct{}"
	}

	return strings.Join([]string{"RestoreHbaInfoResponse", string(data)}, " ")
}
