package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkRequest Request Object
type ShowNetworkRequest struct {

	// **参数解释**：网络的ID，取值自网络详情的metadata.name字段。 **约束限制**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为[36-63]个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NetworkName string `json:"network_name"`
}

func (o ShowNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkRequest struct{}"
	}

	return strings.Join([]string{"ShowNetworkRequest", string(data)}, " ")
}
