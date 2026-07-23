package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoQueryShareDocs struct {

	// **参数解释：**  应用ID，用于指定实例所属的应用。 当URL路径中已能明确应用上下文时，此参数可不传。 获取方法请参见[获取租户下的应用清单 - ListApps](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  - 于2023年06月01日之前创建的应用：由英文字母和数字组成，长度为1-36个字符。 - 于2023年06月01日之后创建的应用：由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`

	// **参数解释：**  结构化文档ID，系统生成的文档主键唯一标识，用于指定待查询分享授权列表的目标文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params string `json:"params"`
}

func (o RdmParamVoQueryShareDocs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoQueryShareDocs struct{}"
	}

	return strings.Join([]string{"RdmParamVoQueryShareDocs", string(data)}, " ")
}
