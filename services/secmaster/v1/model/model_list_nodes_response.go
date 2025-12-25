package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesResponse Response Object
type ListNodesResponse struct {

	// 节点数量
	Count *int64 `json:"count,omitempty"`

	// 节点
	Records        *[]Node `json:"records,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
