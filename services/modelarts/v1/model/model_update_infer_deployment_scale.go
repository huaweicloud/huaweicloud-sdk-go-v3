package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInferDeploymentScale struct {

	// **参数解释：** 服务实例数。 **取值范围：** [1, 128]。
	Count int32 `json:"count"`
}

func (o UpdateInferDeploymentScale) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInferDeploymentScale struct{}"
	}

	return strings.Join([]string{"UpdateInferDeploymentScale", string(data)}, " ")
}
