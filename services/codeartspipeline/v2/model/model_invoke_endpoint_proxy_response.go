package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeEndpointProxyResponse Response Object
type InvokeEndpointProxyResponse struct {
	Error *ErrorInfo `json:"error,omitempty"`

	// **参数解释**： 错误结果。 **取值范围**： 不涉及。
	Result *interface{} `json:"result,omitempty"`

	// **参数解释**： 状态值。 **取值范围**： 不涉及。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InvokeEndpointProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeEndpointProxyResponse struct{}"
	}

	return strings.Join([]string{"InvokeEndpointProxyResponse", string(data)}, " ")
}
