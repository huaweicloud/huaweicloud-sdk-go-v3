package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateResourceShareTagsRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *TagResourceReqBody `json:"body,omitempty"`
}

func (o BatchCreateResourceShareTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceShareTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceShareTagsRequest", string(data)}, " ")
}
