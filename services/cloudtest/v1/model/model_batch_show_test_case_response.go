package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowTestCaseResponse Response Object
type BatchShowTestCaseResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchShowTestCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowTestCaseResponse struct{}"
	}

	return strings.Join([]string{"BatchShowTestCaseResponse", string(data)}, " ")
}
