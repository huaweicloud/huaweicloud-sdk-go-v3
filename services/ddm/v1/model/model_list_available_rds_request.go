package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableRdsRequest Request Object
type ListAvailableRdsRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListAvailableRdsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsRequest", string(data)}, " ")
}
