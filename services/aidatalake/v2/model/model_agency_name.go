package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyName **参数解释**：委托名称。 **约束限制**：具体请参见[[配置Aura委托](https://support.huaweicloud.com/usermanual-aidatalake/aidatalake_01_0078.html#section4)](tag:HEC)[[配置Aura委托](https://support.huaweicloud.com/intl/zh-cn/usermanual-aidatalake/aidatalake_01_0078.html#section4)](tag:hws_hk)。 **取值范围**：不涉及。 **默认取值**：不涉及。
type AgencyName struct {
}

func (o AgencyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyName struct{}"
	}

	return strings.Join([]string{"AgencyName", string(data)}, " ")
}
