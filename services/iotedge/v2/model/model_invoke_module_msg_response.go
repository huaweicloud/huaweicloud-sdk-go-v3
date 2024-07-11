package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeModuleMsgResponse Response Object
type InvokeModuleMsgResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o InvokeModuleMsgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeModuleMsgResponse struct{}"
	}

	return strings.Join([]string{"InvokeModuleMsgResponse", string(data)}, " ")
}
