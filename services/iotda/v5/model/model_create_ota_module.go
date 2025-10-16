package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOtaModule 添OTA模块对象结构体
type CreateOtaModule struct {

	// **参数说明**：资源空间ID。存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的升级包归属到哪个资源空间下。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	AppId string `json:"app_id"`

	// **参数说明**：设备关联的产品ID，用于唯一标识一个产品模型，创建产品后获得。方法请参见 [[创建产品](https://support.huaweicloud.com/api-iothub/iot_06_v5_0050.html)](tag:hws)[[创建产品](https://support.huaweicloud.com/intl/zh-cn/api-iothub/iot_06_v5_0050.html)](tag:hws_hk)。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	ProductId string `json:"product_id"`

	// **参数说明**：OTA模块名称，产品下唯一且不可修改。 **取值范围**：长度不超过64，只允许英文字母、数字、下划线（_）、连接符（-）、英文点（.）的组合。
	ModuleName string `json:"module_name"`

	// **参数说明**：OTA模块别名。 **取值范围**：长度不超过64，只允许中文、英文字母、数字、下划线（_）、连接符（-）、英文点（.）的组合。
	AliasName *string `json:"alias_name,omitempty"`

	// **参数说明**：用于描述模块的功能等信息。 **取值范围**：长度不超过1024。
	Description *string `json:"description,omitempty"`
}

func (o CreateOtaModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOtaModule struct{}"
	}

	return strings.Join([]string{"CreateOtaModule", string(data)}, " ")
}
