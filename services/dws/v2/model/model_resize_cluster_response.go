package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterResponse Response Object
type ResizeClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResizeClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterResponse struct{}"
	}

	return strings.Join([]string{"ResizeClusterResponse", string(data)}, " ")
}
