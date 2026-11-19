package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaurusDbSlowProcessThresholdRequest Request Object
type ShowTaurusDbSlowProcessThresholdRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  通过调用[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)接口获取实例ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`
}

func (o ShowTaurusDbSlowProcessThresholdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaurusDbSlowProcessThresholdRequest struct{}"
	}

	return strings.Join([]string{"ShowTaurusDbSlowProcessThresholdRequest", string(data)}, " ")
}
