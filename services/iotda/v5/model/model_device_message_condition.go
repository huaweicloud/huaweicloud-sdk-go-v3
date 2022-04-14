package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 条件中设备消息类型的信息，自定义结构。
type DeviceMessageCondition struct {
	// **参数说明**：设备关联的产品ID，用于唯一标识一个产品模型，创建产品后获得。方法请参见 [创建产品](https://support.huaweicloud.com/api-iothub/iot_06_v5_0050.html)。

	ProductId *string `json:"product_id,omitempty"`
	// **参数说明**：产品关联的topic信息，用于过滤消息中指定topic消息。

	Topic *string `json:"topic,omitempty"`
}

func (o DeviceMessageCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceMessageCondition struct{}"
	}

	return strings.Join([]string{"DeviceMessageCondition", string(data)}, " ")
}
