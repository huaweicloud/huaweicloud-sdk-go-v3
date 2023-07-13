package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesTagsRequest Request Object
type ListImagesTagsRequest struct {
}

func (o ListImagesTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImagesTagsRequest", string(data)}, " ")
}
