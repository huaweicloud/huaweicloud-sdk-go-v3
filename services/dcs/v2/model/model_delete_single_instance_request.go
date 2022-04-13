package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSingleInstanceRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o DeleteSingleInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSingleInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteSingleInstanceRequest", string(data)}, " ")
}
