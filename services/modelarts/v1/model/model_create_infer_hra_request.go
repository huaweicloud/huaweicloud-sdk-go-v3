package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferHraRequest Request Object
type CreateInferHraRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`

	// **参数解释：** 部署ID，在[添加部署](CreateInferDeployment.xml)时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其中deployment_id字段即为部署ID。 **约束限制：** 不涉及。 **取值范围：** 部署ID。 **默认取值：** 不涉及。
	DeploymentId string `json:"deployment_id"`

	Body *CreateHraRequestBody `json:"body,omitempty"`
}

func (o CreateInferHraRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferHraRequest struct{}"
	}

	return strings.Join([]string{"CreateInferHraRequest", string(data)}, " ")
}
