package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSubNetworkInterfacesByTagsRequest Request Object
type CountSubNetworkInterfacesByTagsRequest struct {
	Body *CountSubNetworkInterfacesByTagsRequestBody `json:"body,omitempty"`
}

func (o CountSubNetworkInterfacesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSubNetworkInterfacesByTagsRequest struct{}"
	}

	return strings.Join([]string{"CountSubNetworkInterfacesByTagsRequest", string(data)}, " ")
}
