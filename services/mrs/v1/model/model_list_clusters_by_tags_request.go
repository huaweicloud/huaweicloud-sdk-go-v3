package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersByTagsRequest Request Object
type ListClustersByTagsRequest struct {
	Body *ListResourceReq `json:"body,omitempty"`
}

func (o ListClustersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClustersByTagsRequest", string(data)}, " ")
}
