package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppInstanceRequest Request Object
type DeleteAppInstanceRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	// 应用实例ID
	AppInstanceId string `json:"app_instance_id"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o DeleteAppInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppInstanceRequest", string(data)}, " ")
}
