package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVgwConnectionsResponse Response Object
type ListP2cVgwConnectionsResponse struct {
	Connections *[]Connection `json:"connections,omitempty"`

	// 总计数量
	TotalCount *int64 `json:"total_count,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListP2cVgwConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListP2cVgwConnectionsResponse", string(data)}, " ")
}
