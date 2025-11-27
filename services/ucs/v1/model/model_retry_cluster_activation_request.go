package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryClusterActivationRequest Request Object
type RetryClusterActivationRequest struct {

	// 集群ID
	Clusterid string `json:"clusterid"`
}

func (o RetryClusterActivationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryClusterActivationRequest struct{}"
	}

	return strings.Join([]string{"RetryClusterActivationRequest", string(data)}, " ")
}
