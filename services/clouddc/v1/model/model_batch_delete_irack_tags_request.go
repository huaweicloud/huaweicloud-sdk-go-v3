package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIrackTagsRequest Request Object
type BatchDeleteIrackTagsRequest struct {

	// 资源id
	Id string `json:"id"`

	Body *ResourceTags `json:"body,omitempty"`
}

func (o BatchDeleteIrackTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIrackTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteIrackTagsRequest", string(data)}, " ")
}
