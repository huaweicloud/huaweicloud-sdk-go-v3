package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppInstanceRequest Request Object
type UpdateAppInstanceRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	// 应用实例ID
	AppInstanceId string `json:"app_instance_id"`

	Body *UpdateAppInstanceRequestDto `json:"body,omitempty"`
}

func (o UpdateAppInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppInstanceRequest", string(data)}, " ")
}
