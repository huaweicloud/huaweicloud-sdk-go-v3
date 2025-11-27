package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteDocumentRequestBodyParameters struct {

	// 参数的key
	Key *string `json:"key,omitempty"`

	// 参数的value
	Value *string `json:"value,omitempty"`
}

func (o ExecuteDocumentRequestBodyParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequestBodyParameters struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequestBodyParameters", string(data)}, " ")
}
