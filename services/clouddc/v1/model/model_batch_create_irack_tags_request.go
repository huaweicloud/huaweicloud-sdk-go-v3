package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIrackTagsRequest Request Object
type BatchCreateIrackTagsRequest struct {

	// 资源id
	Id string `json:"id"`

	Body *ResourceTags `json:"body,omitempty"`
}

func (o BatchCreateIrackTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIrackTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateIrackTagsRequest", string(data)}, " ")
}
