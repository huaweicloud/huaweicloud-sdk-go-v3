package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNewHostsResponse Response Object
type ListNewHostsResponse struct {

	// 主机数量
	Total *int32 `json:"total,omitempty"`

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机信息列表
	Result         *[]HostInfo `json:"result,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListNewHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNewHostsResponse struct{}"
	}

	return strings.Join([]string{"ListNewHostsResponse", string(data)}, " ")
}
