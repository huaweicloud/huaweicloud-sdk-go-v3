package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRunningNodesResponse Response Object
type ListRunningNodesResponse struct {

	// 运行中节点数量
	Count *int64 `json:"count,omitempty"`

	// 组件配置
	Records        *[]ComponentConfiguration `json:"records,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListRunningNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRunningNodesResponse struct{}"
	}

	return strings.Join([]string{"ListRunningNodesResponse", string(data)}, " ")
}
