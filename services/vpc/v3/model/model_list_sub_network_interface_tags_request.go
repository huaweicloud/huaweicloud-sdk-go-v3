package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubNetworkInterfaceTagsRequest Request Object
type ListSubNetworkInterfaceTagsRequest struct {
}

func (o ListSubNetworkInterfaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfaceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfaceTagsRequest", string(data)}, " ")
}
