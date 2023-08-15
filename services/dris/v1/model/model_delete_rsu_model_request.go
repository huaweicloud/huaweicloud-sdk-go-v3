package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRsuModelRequest Request Object
type DeleteRsuModelRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：RSU型号ID，用于唯一标识一个RSU型号，在平台创建RSU型号后由平台分配获得。  **取值范围**：长度不小于1不超过36，只允许字母、数字、连接符（-）的组合。
	RsuModelId string `json:"rsu_model_id"`
}

func (o DeleteRsuModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRsuModelRequest struct{}"
	}

	return strings.Join([]string{"DeleteRsuModelRequest", string(data)}, " ")
}
