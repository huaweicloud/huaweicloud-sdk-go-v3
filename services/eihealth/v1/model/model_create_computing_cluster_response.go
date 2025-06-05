package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComputingClusterResponse Response Object
type CreateComputingClusterResponse struct {

	// 计算集群ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateComputingClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateComputingClusterResponse", string(data)}, " ")
}
