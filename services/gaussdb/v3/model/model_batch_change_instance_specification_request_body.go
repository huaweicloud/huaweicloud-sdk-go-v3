package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeInstanceSpecificationRequestBody **参数解释**：  批量变更实例规格请求体。  **约束限制**：  不涉及。
type BatchChangeInstanceSpecificationRequestBody struct {

	// **参数解释**：  批量规格变更的实例ID列表。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  实例ID只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceIds []string `json:"instance_ids"`

	// **参数解释**：  规格码。  获取方法请参见[查询数据库规格](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlFlavors.html)中的响应参数\"spec_code\"。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SpecCode string `json:"spec_code"`

	// **参数解释**：  批量变更包年/包月实例规格时可指定，表示是否自动从客户的账户中支付。  **约束限制**：  不涉及。  **取值范围**： - true：自动支付，默认该方式。 - false：手动支付。  **默认取值**：  true。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o BatchChangeInstanceSpecificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeInstanceSpecificationRequestBody struct{}"
	}

	return strings.Join([]string{"BatchChangeInstanceSpecificationRequestBody", string(data)}, " ")
}
