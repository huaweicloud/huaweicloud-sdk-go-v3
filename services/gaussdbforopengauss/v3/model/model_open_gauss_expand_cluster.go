package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CN横向扩容/DN分片扩容时必填
type OpenGaussExpandCluster struct {
	Coordinators *[]OpenGaussCoordinators `json:"coordinators,omitempty"`

	Shard *OpenGaussShard `json:"shard,omitempty"`
}

func (o OpenGaussExpandCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussExpandCluster struct{}"
	}

	return strings.Join([]string{"OpenGaussExpandCluster", string(data)}, " ")
}
