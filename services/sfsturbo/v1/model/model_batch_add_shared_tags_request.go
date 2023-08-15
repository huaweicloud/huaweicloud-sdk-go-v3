package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddSharedTagsRequest Request Object
type BatchAddSharedTagsRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 共享ID
	ShareId string `json:"share_id"`

	Body *BatchAddSharedTagsRequestBody `json:"body,omitempty"`
}

func (o BatchAddSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsRequest", string(data)}, " ")
}
