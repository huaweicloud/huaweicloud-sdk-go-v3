package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentHpaEventsRequest Request Object
type ListInferDeploymentHpaEventsRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`

	// **参数解释：** 部署ID，在[添加部署](CreateInferDeployment.xml)时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其中deployment_id字段即为部署ID。 **约束限制：** 不涉及。 **取值范围：** 部署ID。 **默认取值：** 不涉及。
	DeploymentId string `json:"deployment_id"`

	// **参数解释：** 指定每一页返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表的起始页，从0开始计数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 排序字段，多个字段以\",\"分隔，支持create_at, update_at，默认值update_at。 **约束限制：** 不涉及。 **取值范围：** - lastProbeTime：按照执行记录事件排序。 - status：按状态排序。 **默认取值：** update_at。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释：** 排序方式。 **约束限制：** 不涉及。 **取值范围：** - ASC: 递增排序。 - DESC: 递减排序。 **默认取值：** DESC。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListInferDeploymentHpaEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentHpaEventsRequest struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentHpaEventsRequest", string(data)}, " ")
}
