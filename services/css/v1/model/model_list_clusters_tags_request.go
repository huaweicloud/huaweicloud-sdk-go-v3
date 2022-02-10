package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClustersTagsRequest struct {
}

func (o ListClustersTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClustersTagsRequest", string(data)}, " ")
}
