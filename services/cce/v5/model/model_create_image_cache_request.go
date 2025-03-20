package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageCacheRequest Request Object
type CreateImageCacheRequest struct {
	Body *CreateImageCacheRequestBody `json:"body,omitempty"`
}

func (o CreateImageCacheRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageCacheRequest struct{}"
	}

	return strings.Join([]string{"CreateImageCacheRequest", string(data)}, " ")
}
