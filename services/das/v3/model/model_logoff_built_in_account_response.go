package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffBuiltInAccountResponse Response Object
type LogoffBuiltInAccountResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o LogoffBuiltInAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffBuiltInAccountResponse struct{}"
	}

	return strings.Join([]string{"LogoffBuiltInAccountResponse", string(data)}, " ")
}
