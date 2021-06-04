package model

import (
	"encoding/json"

	"strings"
)

// 条件中设备消息类型的信息，自定义结构。
type DeviceMessageCondition struct {
	// **参数说明**：设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。

	ProductId *string `json:"product_id,omitempty"`
	// **参数说明**：产品关联的topic信息，用于过滤消息中指定topic消息。

	Topic *string `json:"topic,omitempty"`
}

func (o DeviceMessageCondition) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeviceMessageCondition struct{}"
	}

	return strings.Join([]string{"DeviceMessageCondition", string(data)}, " ")
}
