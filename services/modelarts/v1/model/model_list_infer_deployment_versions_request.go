package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentVersionsRequest Request Object
type ListInferDeploymentVersionsRequest struct {

	// **参数解释：** 服务ID，在创建服务时即可在返回体中获取，也可通过查询服务列表接口获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`

	// **参数解释：** 部署ID，在[添加部署](CreateInferDeployment.xml)时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其中deployment_id字段即为部署ID。 **约束限制：** 不涉及。 **取值范围：** 部署ID。 **默认取值：** 不涉及。
	DeploymentId string `json:"deployment_id"`

	// **参数解释：** 排序字段。 **约束限制：** 不涉及。 **取值范围：** - create_at：按创建时间排序。 - update_at：按更新时间排序。 **默认取值：** update_at。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInferDeploymentVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentVersionsRequest", string(data)}, " ")
}
