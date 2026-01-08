package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolsResponse Response Object
type BatchDeletePoolsResponse struct {

	// 后端服务器组批量删除后的响应结果。
	Pools          *[]BatchDeletePoolsResp `json:"pools,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchDeletePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolsResponse", string(data)}, " ")
}
