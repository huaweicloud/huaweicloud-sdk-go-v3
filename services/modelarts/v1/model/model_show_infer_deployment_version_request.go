package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInferDeploymentVersionRequest Request Object
type ShowInferDeploymentVersionRequest struct {

	// **参数解释：** 服务ID，在创建服务时即可在返回体中获取，也可通过查询服务列表接口获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`

	// 版本
	Version string `json:"version"`

	// 部署id
	DeploymentId string `json:"deployment_id"`
}

func (o ShowInferDeploymentVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferDeploymentVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowInferDeploymentVersionRequest", string(data)}, " ")
}
