package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FunctionRequestDto 编解码函数信息
type FunctionRequestDto struct {

	// **参数说明**：资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，可以携带该参数查询指定资源空间下的产品列表，不携带该参数则会查询该用户下所有产品列表。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	AppId *string `json:"app_id,omitempty"`

	// **参数说明**：设备关联的产品ID，用于唯一标识一个产品模型，创建产品后获得。方法请参见 [[创建产品](https://support.huaweicloud.com/api-iothub/iot_06_v5_0050.html)](tag:hws)[[创建产品](https://support.huaweicloud.com/intl/zh-cn/api-iothub/iot_06_v5_0050.html)](tag:hws_hk)。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	ProductId string `json:"product_id"`

	// **参数说明**：产品关联函数的URN（Uniform Resource Name）。 **取值范围**：长度不超过256，只允许字母、数字、下划线（_）、连接符（-）、分隔符（:）的组合。
	Urn string `json:"urn"`

	// **参数说明**：编解码函数描述信息。 **取值范围**：128，只允许中文、字母、数字、以及_? '#().,&%@!-等字符的组合。
	Description *string `json:"description,omitempty"`
}

func (o FunctionRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionRequestDto struct{}"
	}

	return strings.Join([]string{"FunctionRequestDto", string(data)}, " ")
}
