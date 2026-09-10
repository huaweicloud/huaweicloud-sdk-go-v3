package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubNetworkInterfaceTagsResponse Response Object
type ListSubNetworkInterfaceTagsResponse struct {

	// tag对象列表
	Tags *[]ListTag `json:"tags,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 资源数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubNetworkInterfaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfaceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfaceTagsResponse", string(data)}, " ")
}
