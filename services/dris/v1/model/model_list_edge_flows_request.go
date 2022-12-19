package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEdgeFlowsRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：查询此时间后达到平台的消息。  格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z'。  例如 2020-09-01T01:37:01.000Z。  **取值范围**：携带edge_id参数查询时，from_date和to_date的时间范围不能超过24小时；未携带edge_id参数查询时，from_date和to_date的时间范围不能超过1小时。
	FromDate *string `json:"from_date,omitempty"`

	// **参数说明**：查询此时间前达到平台的消息。  格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z'。  例如 2020-09-02T01:37:01.000Z。  **取值范围**：携带edge_id参数查询时，from_date和to_date的时间范围不能超过24小时；未携带edge_id参数查询时，from_date和to_date的时间范围不能超过1小时。
	ToDate *string `json:"to_date,omitempty"`

	// **参数说明**：Edge ID，用于唯一标识一个Edge。  **取值范围**：数字，a至f的小写字母，横杠（-），长度为36的组合。
	EdgeId *string `json:"edge_id,omitempty"`
}

func (o ListEdgeFlowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeFlowsRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeFlowsRequest", string(data)}, " ")
}
