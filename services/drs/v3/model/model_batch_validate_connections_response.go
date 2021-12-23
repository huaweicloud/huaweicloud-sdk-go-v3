package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchValidateConnectionsResponse struct {
	// 批量测试连接响应体集合。

	Results *[]CheckJobResp `json:"results,omitempty"`
	// 总记录数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchValidateConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateConnectionsResponse struct{}"
	}

	return strings.Join([]string{"BatchValidateConnectionsResponse", string(data)}, " ")
}
