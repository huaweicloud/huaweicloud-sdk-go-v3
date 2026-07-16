package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInferDeploymentScaleRequest Request Object
type UpdateInferDeploymentScaleRequest struct {

	// **参数解释：** 服务ID
	ServiceId string `json:"service_id"`

	// **参数解释：** 部署名称，在创建部署时即可在返回体中获取，也可通过查询服务部署列表获取当前用户拥有的部署，其name字段即为部署名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DeploymentName string `json:"deployment_name"`

	Body *UpdateInferDeploymentScale `json:"body,omitempty"`
}

func (o UpdateInferDeploymentScaleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInferDeploymentScaleRequest struct{}"
	}

	return strings.Join([]string{"UpdateInferDeploymentScaleRequest", string(data)}, " ")
}
