package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 根据产品名称和实例，批量修改子设备协议配置请求结构体。
type BatchUpdateConfigs struct {
	// 设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。

	ProductId string `json:"product_id"`
	// **参数说明**：资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定归属到哪个资源空间下，否则将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。

	AppId *string `json:"app_id,omitempty"`
	// 设备协议配置数据。

	ProtocolMapping *interface{} `json:"protocol_mapping"`
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"instance_id,omitempty"`
}

func (o BatchUpdateConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConfigs struct{}"
	}

	return strings.Join([]string{"BatchUpdateConfigs", string(data)}, " ")
}
