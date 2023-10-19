package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkTagsRequest Request Object
type ListCentralNetworkTagsRequest struct {
}

func (o ListCentralNetworkTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkTagsRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkTagsRequest", string(data)}, " ")
}
