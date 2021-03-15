package model

import (
	"encoding/json"

	"strings"
)

// 条件中设备数据类型的信息，自定义结构。
type DeviceDataCondition struct {
	// 设备ID，用于唯一标识一个设备，在注册设备时由物联网平台分配获得。当rule_type为DEVICE_LINKAGE时，该参数值和product_id不能同时为空。如果该参数和product_id同时存在时，以该参数值对应的设备进行条件过滤。
	DeviceId *string `json:"device_id,omitempty"`
	// 设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。当rule_type为DEVICE_LINKAGE时，该参数值和device_id不能同时为空。如果该参数和device_id同时存在时，以device_id参数值对应的设备进行条件过滤。
	ProductId *string `json:"product_id,omitempty"`
	// 数据过滤条件
	Filters *[]PropertyFilter `json:"filters,omitempty"`
}

func (o DeviceDataCondition) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeviceDataCondition struct{}"
	}

	return strings.Join([]string{"DeviceDataCondition", string(data)}, " ")
}
