package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubNetworkInterfaceTagsResponse Response Object
type ShowSubNetworkInterfaceTagsResponse struct {

	// tag对象列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubNetworkInterfaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfaceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfaceTagsResponse", string(data)}, " ")
}
