package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPoolsResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`
	// 后端服务器组列表。

	Pools          *[]Pool `json:"pools,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListPoolsResponse", string(data)}, " ")
}
