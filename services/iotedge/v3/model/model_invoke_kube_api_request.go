package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeKubeApiRequest Request Object
type InvokeKubeApiRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	// 透传的k8s的API，{method} {uri}?{query_param}
	XForwardTarget string `json:"X-Forward-Target"`

	// 透传的API的header
	XForwardHeaders string `json:"X-Forward-Headers"`
}

func (o InvokeKubeApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeKubeApiRequest struct{}"
	}

	return strings.Join([]string{"InvokeKubeApiRequest", string(data)}, " ")
}
