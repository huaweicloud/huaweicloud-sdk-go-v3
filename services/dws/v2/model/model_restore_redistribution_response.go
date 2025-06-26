package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRedistributionResponse Response Object
type RestoreRedistributionResponse struct {

	// **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RestoreRedistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedistributionResponse struct{}"
	}

	return strings.Join([]string{"RestoreRedistributionResponse", string(data)}, " ")
}
