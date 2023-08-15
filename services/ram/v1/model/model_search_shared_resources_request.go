package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSharedResourcesRequest Request Object
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
