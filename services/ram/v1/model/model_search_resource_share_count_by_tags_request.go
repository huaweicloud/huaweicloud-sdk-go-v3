package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceShareCountByTagsRequest Request Object
type SearchResourceShareCountByTagsRequest struct {
	Body *ResourceSharesByTagsReqBody `json:"body,omitempty"`
}

func (o SearchResourceShareCountByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareCountByTagsRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceShareCountByTagsRequest", string(data)}, " ")
}
