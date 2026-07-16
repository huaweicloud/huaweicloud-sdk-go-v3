package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentVersionsResponse Response Object
type ListInferDeploymentVersionsResponse struct {

	// **参数解释：** 在线服务部署数据。
	Data *[]InferDeploymentVersionItemResp `json:"data,omitempty"`

	// **参数解释：** 当前页码，从0开始计数。 **取值范围：** 不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释：** 当前页数量。 **取值范围：** 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释：** 当前页数量。 **取值范围：** 不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释：** 总记录条数。 **取值范围：** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInferDeploymentVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentVersionsResponse", string(data)}, " ")
}
