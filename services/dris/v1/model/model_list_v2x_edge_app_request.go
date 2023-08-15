package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2xEdgeAppRequest Request Object
type ListV2xEdgeAppRequest struct {

	// **参数说明**：Edge ID，用于唯一标识一个Edge，创建Edge后获得。方法参见 [创建Edge](https://support.huaweicloud.com/api-v2x/v2x_04_0078.html)。
	V2xEdgeId string `json:"v2x_edge_id"`

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：每页记录数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListV2xEdgeAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2xEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"ListV2xEdgeAppRequest", string(data)}, " ")
}
