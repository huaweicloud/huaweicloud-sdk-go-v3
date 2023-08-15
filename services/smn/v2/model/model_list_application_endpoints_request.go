package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationEndpointsRequest Request Object
type ListApplicationEndpointsRequest struct {

	// Application的唯一资源标识，可通过[查询Application](smn_api_57004.xml)获取该标识。
	ApplicationUrn string `json:"application_urn"`

	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 设备是否可用，值为true或false字符串。
	Enabled *string `json:"enabled,omitempty"`

	// 设备token，最大长度512个字节。
	Token *string `json:"token,omitempty"`

	// 用户数据，最大长度2048个字节。
	UserData *string `json:"user_data,omitempty"`
}

func (o ListApplicationEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointsRequest", string(data)}, " ")
}
