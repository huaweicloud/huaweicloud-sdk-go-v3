package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowRadarsRequest Request Object
type BatchShowRadarsRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：设备序列号。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - UNKNOWN：未知 - SLEEP：睡眠
	Status *string `json:"status,omitempty"`
}

func (o BatchShowRadarsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowRadarsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowRadarsRequest", string(data)}, " ")
}
