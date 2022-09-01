package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApplicationEndpointsRequest struct {

	// Application的唯一资源标识，可通过[查询Application](https://support.huaweicloud.com/api-smn/ListApplications.html)获取该标识。
	ApplicationUrn string `json:"application_urn" xml:"application_urn"`

	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询的数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 设备是否可用，值为true或false字符串。
	Enabled *string `json:"enabled,omitempty" xml:"enabled"`

	// 设备token，最大长度512个字节。
	Token *string `json:"token,omitempty" xml:"token"`

	// 用户数据，最大长度2048个字节。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`
}

func (o ListApplicationEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointsRequest", string(data)}, " ")
}
