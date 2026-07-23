package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListVersionModelVersionUndoCheckOutDto struct {

	// **参数解释：**  应用ID，用于指定实例所属的应用。 当URL路径中已能明确应用上下文时，此参数可不传。 获取方法请参见[获取租户下的应用清单 - ListApps](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  由英文字母和数字组成，且长度固定为32个字符。  **默认取值：**  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	// **参数解释：**  M-V模型批量撤销检出请求参数数组，每个元素为一个撤销检出对象。  **约束限制：**  单次请求不超过1000个。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params []VersionModelVersionUndoCheckOutDto `json:"params"`
}

func (o RdmParamVoListVersionModelVersionUndoCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListVersionModelVersionUndoCheckOutDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListVersionModelVersionUndoCheckOutDto", string(data)}, " ")
}
