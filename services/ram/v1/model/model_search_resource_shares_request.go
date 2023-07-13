package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceSharesRequest Request Object
type SearchResourceSharesRequest struct {
	Body *SearchResourceSharesReqBody `json:"body,omitempty"`
}

func (o SearchResourceSharesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceSharesRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceSharesRequest", string(data)}, " ")
}
