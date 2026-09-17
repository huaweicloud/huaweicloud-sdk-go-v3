package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetModelConfigResponse Response Object
type GetModelConfigResponse struct {

	// **参数解释**： 响应状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	Result *ModelConfigDto `json:"result,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetModelConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetModelConfigResponse struct{}"
	}

	return strings.Join([]string{"GetModelConfigResponse", string(data)}, " ")
}
