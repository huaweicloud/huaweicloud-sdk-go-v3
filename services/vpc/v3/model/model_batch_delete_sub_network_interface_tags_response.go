package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubNetworkInterfaceTagsResponse Response Object
type BatchDeleteSubNetworkInterfaceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSubNetworkInterfaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubNetworkInterfaceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubNetworkInterfaceTagsResponse", string(data)}, " ")
}
