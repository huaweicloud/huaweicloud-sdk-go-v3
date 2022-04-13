package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateClusterRequestBody struct {
	Cluster *CreateClusterInfo `json:"cluster"`
}

func (o CreateClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateClusterRequestBody", string(data)}, " ")
}
