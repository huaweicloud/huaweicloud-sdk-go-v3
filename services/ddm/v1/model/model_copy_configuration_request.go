package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyConfigurationRequest Request Object
type CopyConfigurationRequest struct {

	// **参数解释**：  参数组ID。不能使用默认参数组进行复制  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	ConfigId string `json:"config_id"`

	Body *ConfigurationCopyReqV3 `json:"body,omitempty"`
}

func (o CopyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequest", string(data)}, " ")
}
