package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRedistributionPriorityResponse Response Object
type SetRedistributionPriorityResponse struct {

	// **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SetRedistributionPriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRedistributionPriorityResponse struct{}"
	}

	return strings.Join([]string{"SetRedistributionPriorityResponse", string(data)}, " ")
}
