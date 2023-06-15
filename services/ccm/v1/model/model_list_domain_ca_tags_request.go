package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDomainCaTagsRequest struct {
}

func (o ListDomainCaTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainCaTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainCaTagsRequest", string(data)}, " ")
}
