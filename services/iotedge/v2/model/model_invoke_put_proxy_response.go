package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokePutProxyResponse Response Object
type InvokePutProxyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o InvokePutProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokePutProxyResponse struct{}"
	}

	return strings.Join([]string{"InvokePutProxyResponse", string(data)}, " ")
}
