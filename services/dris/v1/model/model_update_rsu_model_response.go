package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateRsuModelResponse struct {

	// **参数说明**：RSU型号ID，用于唯一标识一个RSU型号，在平台创建RSU型号后由平台分配获得。  **取值范围**：长度不小于1不超过36，只允许字母、数字、连接符（-）的组合。
	RsuModelId *string `json:"rsu_model_id,omitempty"`

	// **参数说明**: RSU型号名称。  **取值范围**：长度不低于1不超过64，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**: RSU的厂商名称。  **取值范围**：长度不低于1不超过32，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	ManufacturerName *string `json:"manufacturer_name,omitempty"`

	// **参数说明**: RSU型号的描述信息。  **取值范围**：长度不低于1不超过128，只允许中文、字母、数字、下划线（_）、问号（?）、反引号（'）、井号（#）、左小括号（(）、右小括号（)）、点（.）、逗号（,）、与（&）、百分号（%）、At（@）、感叹号（!）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**: RSU型号更新的时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'  例如：2020-12-07T01:32:17Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`

	// **参数说明**: 在平台创建RSU型号的时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'  例如：2020-12-07T01:32:17Z
	CreatedTime    *sdktime.SdkTime `json:"created_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateRsuModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRsuModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateRsuModelResponse", string(data)}, " ")
}
