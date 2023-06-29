package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppInstanceHistoryRequest Request Object
type ListAppInstanceHistoryRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	// 应用实例ID
	AppInstanceId string `json:"app_instance_id"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o ListAppInstanceHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppInstanceHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListAppInstanceHistoryRequest", string(data)}, " ")
}
