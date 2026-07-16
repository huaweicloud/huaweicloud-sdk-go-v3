package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchInferDeploymentVersionRequestBody struct {

	// **参数解释：** 服务ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** 待切换的目标版本。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TargetDeploymentVersion *string `json:"target_deployment_version,omitempty"`

	// **参数解释：** 部署ID。
	InferName string `json:"infer_name"`
}

func (o SwitchInferDeploymentVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchInferDeploymentVersionRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchInferDeploymentVersionRequestBody", string(data)}, " ")
}
