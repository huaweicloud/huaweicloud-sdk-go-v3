package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIpGroupResponse struct {

	// 该用户总的Ip地址组数量，包含本地与共享地址组
	CloudTotal *int32 `json:"cloudTotal,omitempty"`

	// 该用户当前企业项目下Ip地址组数量，只包含本地地址组
	Total *int32 `json:"total,omitempty"`

	// 地址组信息列表
	Items          *[]IpGroupBody `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ListIpGroupResponse", string(data)}, " ")
}
