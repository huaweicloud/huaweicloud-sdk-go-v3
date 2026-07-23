package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListShareDocsParamDto struct {

	// **参数解释：**  应用ID，用于指定实例所属的应用。 当URL路径中已能明确应用上下文时，此参数可不传。 获取方法请参见[获取租户下的应用清单 - ListApps](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  - 于2023年06月01日之前创建的应用：由英文字母和数字组成，长度为1-36个字符。 - 于2023年06月01日之后创建的应用：由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	// **参数解释：**  批量分享参数数组，每个元素包含一条文档分享配置信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params []ShareDocsParamDto `json:"params"`
}

func (o RdmParamVoListShareDocsParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListShareDocsParamDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListShareDocsParamDto", string(data)}, " ")
}
