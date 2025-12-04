package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartDdmInstanceRequest Request Object
type RestartDdmInstanceRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`
}

func (o RestartDdmInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartDdmInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartDdmInstanceRequest", string(data)}, " ")
}
