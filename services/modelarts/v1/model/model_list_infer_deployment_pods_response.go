package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentPodsResponse Response Object
type ListInferDeploymentPodsResponse struct {

	// **参数解释：** 当前页码，从0开始计数。 **取值范围：** 不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释：** 服务的pod的列表。
	Data *[]ServicePodResponse `json:"data,omitempty"`

	// **参数解释：** 总页数，根据传入的limit字段和数据总条数计算得出。如总记录条数为10，limit（单页最大条目数）为3，则页数为4。 **取值范围：** 不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释：** 当前页数量。 **取值范围：** 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释：** 总记录条数。 **取值范围：** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInferDeploymentPodsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentPodsResponse struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentPodsResponse", string(data)}, " ")
}
