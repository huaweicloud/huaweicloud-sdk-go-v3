package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteInstancesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o DeleteInstancesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesV2Request struct{}"
	}

	return strings.Join([]string{"DeleteInstancesV2Request", string(data)}, " ")
}
