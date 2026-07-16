package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferDeploymentInstanceRequest Request Object
type DeleteInferDeploymentInstanceRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 部署名称。
	DeploymentName string `json:"deployment_name"`

	// **参数解释：** 服务实例名字，可以为all，为all时去查询所有的服务实例。 **约束限制：** 不涉及。 **取值范围：** 服务实例名字。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 是否强制删除。 **约束限制：** 不涉及。 **取值范围：** - true：强制删除。 - false：不强制删除。 **默认取值：** false。
	Force *bool `json:"force,omitempty"`

	// **参数解释：** 删除操作类型。 **约束限制：** 枚举值。 **取值范围：** - DELETE：直接删除，释放资源。 - RECREATE：删除后重建。 **默认取值：** RECREATE。
	Operation *string `json:"operation,omitempty"`
}

func (o DeleteInferDeploymentInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferDeploymentInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInferDeploymentInstanceRequest", string(data)}, " ")
}
