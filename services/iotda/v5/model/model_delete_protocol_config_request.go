package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtocolConfigRequest Request Object
type DeleteProtocolConfigRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，建议携带该参数，在使用专业版时必须携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID，具体获取方式请参考[[查看实例详情](https://support.huaweicloud.com/usermanual-iothub/iot_01_0079.html#section1)](tag:hws) [[查看实例详情](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0079.html#section1)](tag:hws_hk)。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：泛协议配置ID。
	ProtocolId string `json:"protocol_id"`
}

func (o DeleteProtocolConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtocolConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtocolConfigRequest", string(data)}, " ")
}
