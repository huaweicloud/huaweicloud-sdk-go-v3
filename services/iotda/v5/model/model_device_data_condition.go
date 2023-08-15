package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceDataCondition 条件中设备数据类型的信息，自定义结构。
type DeviceDataCondition struct {

	// **参数说明**：设备ID，用于唯一标识一个设备，在注册设备时由物联网平台分配获得。存在该参数时设备属性触发根据指定设备触发，该参数值和product_id不能同时为空。如果该参数和product_id同时存在时，以该参数值对应的设备进行条件过滤。 **取值范围**：长度不超过128，只允许字母、数字、下划线（_）、连接符（-）的组合。
	DeviceId *string `json:"device_id,omitempty"`

	// **参数说明**：设备关联的产品ID，用于唯一标识一个产品模型，创建产品后获得。方法请参见 [[创建产品](https://support.huaweicloud.com/api-iothub/iot_06_v5_0050.html)](tag:hws)[[创建产品](https://support.huaweicloud.com/intl/zh-cn/api-iothub/iot_06_v5_0050.html)](tag:hws_hk)。存在该参数且device_id为空时设备属性触发匹配该产品下所有设备触发，该参数值和device_id不能同时为空。
	ProductId *string `json:"product_id,omitempty"`

	// 数据过滤条件。
	Filters *[]PropertyFilter `json:"filters,omitempty"`
}

func (o DeviceDataCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceDataCondition struct{}"
	}

	return strings.Join([]string{"DeviceDataCondition", string(data)}, " ")
}
