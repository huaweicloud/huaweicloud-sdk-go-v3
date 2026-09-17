package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokePostProxyResponse Response Object
type InvokePostProxyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o InvokePostProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokePostProxyResponse struct{}"
	}

	return strings.Join([]string{"InvokePostProxyResponse", string(data)}, " ")
}
