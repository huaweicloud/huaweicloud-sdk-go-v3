package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeploymentPatchRequest 更新应用部署部分参数
type DeploymentPatchRequest struct {

	// 技能版本，可选，当下发的技能版本和当前部署的版本不一致的时候，根据技能模板参数更新部署
	Version *string `json:"version,omitempty"`

	Patch *Patch `json:"patch,omitempty"`
}

func (o DeploymentPatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentPatchRequest struct{}"
	}

	return strings.Join([]string{"DeploymentPatchRequest", string(data)}, " ")
}
