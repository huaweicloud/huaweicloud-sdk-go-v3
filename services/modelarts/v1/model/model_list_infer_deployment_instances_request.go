package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentInstancesRequest Request Object
type ListInferDeploymentInstancesRequest struct {

	// **参数解释：** 服务唯一id
	Id string `json:"id"`

	// **参数解释：** 服务部署名字，可以为all
	Name string `json:"name"`

	// **参数解释：** 服务实例状态，一次支持多种状态筛选，多种状态以\",\"连接，不能存在空格。默认不过滤。取值范围有4种RUNNING（运行中）、ERROR（错误）、INIT（初始化）、DELETED（已删除)。 **约束限制：** 不涉及。
	Status *[]string `json:"status,omitempty"`

	// **参数解释：** 指定每一页返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表的起始页。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** pod名字。 **取值范围：** 不涉及。
	PodName *string `json:"pod_name,omitempty"`
}

func (o ListInferDeploymentInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentInstancesRequest", string(data)}, " ")
}
