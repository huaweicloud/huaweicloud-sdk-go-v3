package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteColdTableDataRequest Request Object
type DeleteColdTableDataRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  冷数据表空间ID。  获取方法请参见[查询冷表元信息](https://support.huaweicloud.com/api-taurusdb/ListColdTableMetas.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由数字组成。  **默认取值**：  不涉及。
	SpaceId string `json:"space_id"`

	// **参数解释**：  冷数据表ID。  获取方法请参见[查询冷表元信息](https://support.huaweicloud.com/api-taurusdb/ListColdTableMetas.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由数字组成。  **默认取值**：  不涉及。
	DdId string `json:"dd_id"`
}

func (o DeleteColdTableDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteColdTableDataRequest struct{}"
	}

	return strings.Join([]string{"DeleteColdTableDataRequest", string(data)}, " ")
}
