package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReduceNodeOpenRequest struct {

	// 节点id列表。
	NodeIds []string `json:"node_ids"`
}

func (o ReduceNodeOpenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceNodeOpenRequest struct{}"
	}

	return strings.Join([]string{"ReduceNodeOpenRequest", string(data)}, " ")
}
