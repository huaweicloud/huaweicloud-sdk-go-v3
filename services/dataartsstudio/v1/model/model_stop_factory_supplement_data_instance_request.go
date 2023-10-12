package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFactorySupplementDataInstanceRequest Request Object
type StopFactorySupplementDataInstanceRequest struct {

	// 补数据名称
	InstanceName string `json:"instance_name"`

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`
}

func (o StopFactorySupplementDataInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFactorySupplementDataInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopFactorySupplementDataInstanceRequest", string(data)}, " ")
}
