package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeDeleteProxyResponse Response Object
type InvokeDeleteProxyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o InvokeDeleteProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeDeleteProxyResponse struct{}"
	}

	return strings.Join([]string{"InvokeDeleteProxyResponse", string(data)}, " ")
}
