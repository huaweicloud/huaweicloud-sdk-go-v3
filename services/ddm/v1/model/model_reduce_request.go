package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReduceRequest This is a auto request Object
type ReduceRequest struct {

	// 需要缩容的节点数量，最大值为实例节点数减1。
	NodeNumber int32 `json:"node_number"`

	// 组id，指定当前进行节点扩容的组。当实例的组>1时，必填。
	GroupId *string `json:"group_id,omitempty"`
}

func (o ReduceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceRequest struct{}"
	}

	return strings.Join([]string{"ReduceRequest", string(data)}, " ")
}
