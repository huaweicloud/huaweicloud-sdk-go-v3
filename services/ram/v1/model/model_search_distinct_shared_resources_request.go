package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDistinctSharedResourcesRequest Request Object
type SearchDistinctSharedResourcesRequest struct {
	Body *SearchDistinctSharedResourcesReqBody `json:"body,omitempty"`
}

func (o SearchDistinctSharedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDistinctSharedResourcesRequest struct{}"
	}

	return strings.Join([]string{"SearchDistinctSharedResourcesRequest", string(data)}, " ")
}
