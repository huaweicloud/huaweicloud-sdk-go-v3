package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterKubeconfigRequest Request Object
type CreateClusterKubeconfigRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`
}

func (o CreateClusterKubeconfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterKubeconfigRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterKubeconfigRequest", string(data)}, " ")
}
