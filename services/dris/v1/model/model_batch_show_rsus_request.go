package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowRsusRequest Request Object
type BatchShowRsusRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：RSU的唯一标识符，在平台创建RSU时由平台生成。
	RsuId *string `json:"rsu_id,omitempty"`

	// **参数说明**：RSU的设备序列号。  **取值范围**：只允许字母、数字、下划线（_）的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：设备状态。  **取值范围**：  - ONLINE：在线  - OFFLINE：离线  - INITIAL：初始化  - UNKNOWN：未知
	Status *string `json:"status,omitempty"`

	// **参数说明**：RSU型号ID，用于唯一标识一个RSU型号，在平台创建RSU型号后由平台分配获得，获取方法可参见 [创建RSU型号](https://support.huaweicloud.com/api-v2x/v2x_04_0020.html)。  **取值范围**：长度不低于1不超过36，只允许字母、数字、连接符（-）的组合。  **该字段仅供使用MQTT协议RSU设备的用户输入。使用websocket协议RSU设备的用户需忽略此字段。**
	RsuModelId *string `json:"rsu_model_id,omitempty"`
}

func (o BatchShowRsusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowRsusRequest struct{}"
	}

	return strings.Join([]string{"BatchShowRsusRequest", string(data)}, " ")
}
