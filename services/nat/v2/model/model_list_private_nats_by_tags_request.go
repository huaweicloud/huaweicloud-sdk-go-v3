package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateNatsByTagsRequest Request Object
type ListPrivateNatsByTagsRequest struct {
	Body *ListTagResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListPrivateNatsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsByTagsRequest", string(data)}, " ")
}
