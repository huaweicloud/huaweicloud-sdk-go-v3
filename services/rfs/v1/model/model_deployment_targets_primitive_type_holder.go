package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentTargetsPrimitiveTypeHolder struct {
	DeploymentTargets *DeploymentTargets `json:"deployment_targets"`
}

func (o DeploymentTargetsPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentTargetsPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"DeploymentTargetsPrimitiveTypeHolder", string(data)}, " ")
}
