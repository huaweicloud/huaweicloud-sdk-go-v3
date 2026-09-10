package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubNetworkInterfacesByTagsResponse Response Object
type ListSubNetworkInterfacesByTagsResponse struct {

	// 资源列表
	Resources *[]ListResourceResp `json:"resources,omitempty"`

	// 资源数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubNetworkInterfacesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfacesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfacesByTagsResponse", string(data)}, " ")
}
