package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterConfigRequestBody 多云集群的kubeconfig解析
type MultiCloudClusterConfigRequestBody struct {

	// kubeconfig文件
	KubeConfig string `json:"kube_config"`
}

func (o MultiCloudClusterConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterConfigRequestBody struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterConfigRequestBody", string(data)}, " ")
}
