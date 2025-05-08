package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteDocumentRequsetBodySysTags struct {

	// 系统标签的key
	Key *string `json:"key,omitempty"`

	// 系统标签的value
	Value *string `json:"value,omitempty"`
}

func (o ExecuteDocumentRequsetBodySysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequsetBodySysTags struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequsetBodySysTags", string(data)}, " ")
}
