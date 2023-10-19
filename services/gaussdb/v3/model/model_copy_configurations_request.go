package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyConfigurationsRequest Request Object
type CopyConfigurationsRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数组ID。  通过调用[查询参数模板](https://support.huaweicloud.com/api-gaussdb/ListGaussMySqlConfigurations.html)接口获取。  请求响应成功后在响应消息体中包含的“id”的值即为configuration_id值。
	ConfigurationId string `json:"configuration_id"`

	Body *CopyConfigurationsRequestBody `json:"body,omitempty"`
}

func (o CopyConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"CopyConfigurationsRequest", string(data)}, " ")
}
