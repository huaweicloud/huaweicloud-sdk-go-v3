package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterAllNodeLabelResponse Response Object
type ListClusterAllNodeLabelResponse struct {

	// 节点标签数量
	Count *int64 `json:"count,omitempty"`

	// 数据对象列表
	Labels         *[]NodeLabelRsp `json:"labels,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListClusterAllNodeLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterAllNodeLabelResponse struct{}"
	}

	return strings.Join([]string{"ListClusterAllNodeLabelResponse", string(data)}, " ")
}
