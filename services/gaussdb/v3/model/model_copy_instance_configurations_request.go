package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyInstanceConfigurationsRequest Request Object
type CopyInstanceConfigurationsRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 参数组ID。  通过调用[查询实例详情信息](https://support.huaweicloud.com/api-gaussdbformysql/ShowGaussMySqlInstanceInfo.html)接口获取。  请求响应成功后在响应消息体中包含的“configuration_id”的值即为configuration_id值。
	ConfigurationId string `json:"configuration_id"`

	Body *CopyInstanceConfigurationsRequestBody `json:"body,omitempty"`
}

func (o CopyInstanceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"CopyInstanceConfigurationsRequest", string(data)}, " ")
}
