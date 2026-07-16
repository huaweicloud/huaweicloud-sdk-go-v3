package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDevServerRequest Request Object
type UpdateDevServerRequest struct {

	// **参数解释**：DevServer ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *UpdateServerRequest `json:"body,omitempty"`
}

func (o UpdateDevServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDevServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateDevServerRequest", string(data)}, " ")
}
