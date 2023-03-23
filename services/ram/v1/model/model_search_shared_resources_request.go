package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchSharedResourcesRequest struct {
	Body *SearchSharedResourcesReqBody `json:"body,omitempty"`
}

func (o SearchSharedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSharedResourcesRequest struct{}"
	}

	return strings.Join([]string{"SearchSharedResourcesRequest", string(data)}, " ")
}
