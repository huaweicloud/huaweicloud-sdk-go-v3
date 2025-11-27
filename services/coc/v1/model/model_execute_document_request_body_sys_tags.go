package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteDocumentRequestBodySysTags struct {

	// 系统标签的key
	Key *string `json:"key,omitempty"`

	// 系统标签的value
	Value *string `json:"value,omitempty"`
}

func (o ExecuteDocumentRequestBodySysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequestBodySysTags struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequestBodySysTags", string(data)}, " ")
}
