package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddForwardingConfigsRequest Request Object
type AddForwardingConfigsRequest struct {

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddForwardingConfigRequestDto `json:"body,omitempty"`
}

func (o AddForwardingConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddForwardingConfigsRequest struct{}"
	}

	return strings.Join([]string{"AddForwardingConfigsRequest", string(data)}, " ")
}
