package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentHostsResponse Response Object
type ListEnvironmentHostsResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 总数量
	Total *int32 `json:"total,omitempty"`

	// 环境下主机信息列表
	Result         *[]EnvironmentHostInfo `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListEnvironmentHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentHostsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvironmentHostsResponse", string(data)}, " ")
}
