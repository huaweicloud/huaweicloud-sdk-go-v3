package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterRequestBody This is a auto create Body Object
type CreateLogicalClusterRequestBody struct {
	LogicalCluster *CreateLogicalClusterInfo `json:"logical_cluster"`
}

func (o CreateLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterRequestBody", string(data)}, " ")
}
