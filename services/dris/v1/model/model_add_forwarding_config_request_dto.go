package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddForwardingConfigRequestDto 添加相关资源配置请求对象
type AddForwardingConfigRequestDto struct {

	// **参数说明**：转发配置的类型。  **取值范围**：当前仅支持“kafka，mrskafka”。
	ForwardingType string `json:"forwarding_type"`

	ForwardingConfig *ForwardingConfigRequestDto `json:"forwarding_config"`
}

func (o AddForwardingConfigRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddForwardingConfigRequestDto struct{}"
	}

	return strings.Join([]string{"AddForwardingConfigRequestDto", string(data)}, " ")
}
