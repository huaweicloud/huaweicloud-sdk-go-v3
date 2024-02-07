package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipSegmentDomainTagsRequest Request Object
type ListGeipSegmentDomainTagsRequest struct {
}

func (o ListGeipSegmentDomainTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipSegmentDomainTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGeipSegmentDomainTagsRequest", string(data)}, " ")
}
