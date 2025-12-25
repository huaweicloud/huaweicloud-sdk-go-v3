package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorConnectionsResponse Response Object
type ListCollectorConnectionsResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 相关描述信息
	Records        *[]Connection `json:"records,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCollectorConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorConnectionsResponse", string(data)}, " ")
}
