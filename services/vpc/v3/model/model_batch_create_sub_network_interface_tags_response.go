package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSubNetworkInterfaceTagsResponse Response Object
type BatchCreateSubNetworkInterfaceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSubNetworkInterfaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceTagsResponse", string(data)}, " ")
}
