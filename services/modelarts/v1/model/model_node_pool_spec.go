package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolSpec 节点池创建请求体。
type NodePoolSpec struct {
	Resources *PoolResource `json:"resources"`
}

func (o NodePoolSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolSpec struct{}"
	}

	return strings.Join([]string{"NodePoolSpec", string(data)}, " ")
}
