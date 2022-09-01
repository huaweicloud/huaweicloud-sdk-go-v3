package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CountDevicesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o CountDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountDevicesRequest struct{}"
	}

	return strings.Join([]string{"CountDevicesRequest", string(data)}, " ")
}
