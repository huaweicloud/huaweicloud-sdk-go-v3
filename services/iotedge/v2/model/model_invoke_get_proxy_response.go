package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeGetProxyResponse Response Object
type InvokeGetProxyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o InvokeGetProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeGetProxyResponse struct{}"
	}

	return strings.Join([]string{"InvokeGetProxyResponse", string(data)}, " ")
}
