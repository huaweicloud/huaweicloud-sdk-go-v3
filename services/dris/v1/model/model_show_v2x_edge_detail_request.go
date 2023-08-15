package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2xEdgeDetailRequest Request Object
type ShowV2xEdgeDetailRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：Edge ID，用于唯一标识一个Edge。
	V2xEdgeId string `json:"v2x_edge_id"`
}

func (o ShowV2xEdgeDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2xEdgeDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowV2xEdgeDetailRequest", string(data)}, " ")
}
