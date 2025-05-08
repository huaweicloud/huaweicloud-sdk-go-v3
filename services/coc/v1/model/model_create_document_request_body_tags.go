package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDocumentRequestBodyTags struct {

	// 标签的key
	Key *string `json:"key,omitempty"`

	// 标签的value
	Value *string `json:"value,omitempty"`
}

func (o CreateDocumentRequestBodyTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocumentRequestBodyTags struct{}"
	}

	return strings.Join([]string{"CreateDocumentRequestBodyTags", string(data)}, " ")
}
