package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterWithExistedNodesResponse Response Object
type ResizeClusterWithExistedNodesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResizeClusterWithExistedNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterWithExistedNodesResponse struct{}"
	}

	return strings.Join([]string{"ResizeClusterWithExistedNodesResponse", string(data)}, " ")
}
