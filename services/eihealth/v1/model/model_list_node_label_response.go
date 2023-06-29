package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodeLabelResponse Response Object
type ListNodeLabelResponse struct {

	// 节点标签数量
	Count *int64 `json:"count,omitempty"`

	// 数据对象列表
	Labels         *[]NodeLabelRsp `json:"labels,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListNodeLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeLabelResponse struct{}"
	}

	return strings.Join([]string{"ListNodeLabelResponse", string(data)}, " ")
}
