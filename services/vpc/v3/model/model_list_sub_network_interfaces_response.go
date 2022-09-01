package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubNetworkInterfacesResponse struct {

	// 1、功能说明：请求ID 2、取值范围：标准UUID 3、约束：N/A 4、默认值：N/A 5、权限：N/A
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 1、功能说明：辅助弹性网卡查询对象 2、取值范围：N/A 3、约束：N/A 4、默认值：N/A 5、权限：N/A
	SubNetworkInterfaces *[]SubNetworkInterface `json:"sub_network_interfaces,omitempty" xml:"sub_network_interfaces"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSubNetworkInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfacesResponse", string(data)}, " ")
}
