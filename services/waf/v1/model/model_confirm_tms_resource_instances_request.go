package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmTmsResourceInstancesRequest Request Object
type ConfirmTmsResourceInstancesRequest struct {

	// **参数解释：** 资源类型，目前支持waf-instance，waf **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	ResourceType string `json:"resource_type"`

	Body *TmsResourceInstancesRequest `json:"body,omitempty"`
}

func (o ConfirmTmsResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTmsResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmTmsResourceInstancesRequest", string(data)}, " ")
}
