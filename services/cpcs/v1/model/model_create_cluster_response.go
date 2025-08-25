package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterResponse Response Object
type CreateClusterResponse struct {

	// 集群订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
