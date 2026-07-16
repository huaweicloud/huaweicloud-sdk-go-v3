package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferDeploymentVersionRequest Request Object
type DeleteInferDeploymentVersionRequest struct {

	// **参数解释：** 服务ID，在创建服务时即可在返回体中获取，也可通过查询服务列表接口获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`

	// **参数解释：** 版本号，可通过查询version列表查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Version string `json:"version"`

	// 参数解释： 部署ID，在[添加部署](CreateInferDeployment.xml)时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其中deployment_id字段即为部署ID。 约束限制： 不涉及。 取值范围： 部署ID。 默认取值： 不涉及。
	DeploymentId string `json:"deployment_id"`
}

func (o DeleteInferDeploymentVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferDeploymentVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteInferDeploymentVersionRequest", string(data)}, " ")
}
