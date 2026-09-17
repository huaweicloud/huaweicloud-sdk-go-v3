package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokePatchProxyResponse Response Object
type InvokePatchProxyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o InvokePatchProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokePatchProxyResponse struct{}"
	}

	return strings.Join([]string{"InvokePatchProxyResponse", string(data)}, " ")
}
