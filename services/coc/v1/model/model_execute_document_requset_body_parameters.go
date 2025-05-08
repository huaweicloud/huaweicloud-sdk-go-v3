package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteDocumentRequsetBodyParameters struct {

	// 参数的key
	Key *string `json:"key,omitempty"`

	// 参数的value
	Value *string `json:"value,omitempty"`
}

func (o ExecuteDocumentRequsetBodyParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequsetBodyParameters struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequsetBodyParameters", string(data)}, " ")
}
