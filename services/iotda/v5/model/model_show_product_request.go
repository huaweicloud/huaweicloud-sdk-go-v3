package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProductRequest struct {
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// **参数说明**：产品ID，用于唯一标识一个产品，在物联网平台创建产品后由平台分配获得。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。

	ProductId string `json:"product_id"`
	// **参数说明**：资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，必须携带该参数指定要查询的产品属于哪个资源空间，否则接口会提示错误。如果用户存在多资源空间，同时又不想携带该参数，可以联系华为技术支持对用户数据做资源空间合并。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。

	AppId *string `json:"app_id,omitempty"`
}

func (o ShowProductRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProductRequest struct{}"
	}

	return strings.Join([]string{"ShowProductRequest", string(data)}, " ")
}
