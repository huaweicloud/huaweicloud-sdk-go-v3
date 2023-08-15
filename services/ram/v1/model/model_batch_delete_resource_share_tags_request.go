package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceShareTagsRequest Request Object
type BatchDeleteResourceShareTagsRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *UntagResourceReqBody `json:"body,omitempty"`
}

func (o BatchDeleteResourceShareTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceShareTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceShareTagsRequest", string(data)}, " ")
}
