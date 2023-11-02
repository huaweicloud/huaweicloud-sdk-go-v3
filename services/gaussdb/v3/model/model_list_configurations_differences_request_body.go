package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfigurationsDifferencesRequestBody struct {

	// 需要进行比较的源参数模板ID。  通过调用[查询参数模板](https://support.huaweicloud.com/api-gaussdbformysql/ListGaussMySqlConfigurations.html)接口获取。  请求响应成功后在响应消息体中包含的“id”的值即为source_configuration_id值。
	SourceConfigurationId string `json:"source_configuration_id"`

	// 需要进行比较的目标参数模板ID。  通过调用[查询参数模板](https://support.huaweicloud.com/api-gaussdbformysql/ListGaussMySqlConfigurations.html)接口获取。  请求响应成功后在响应消息体中包含的“id”的值即为target_configuration_id值。
	TargetConfigurationId string `json:"target_configuration_id"`
}

func (o ListConfigurationsDifferencesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsDifferencesRequestBody struct{}"
	}

	return strings.Join([]string{"ListConfigurationsDifferencesRequestBody", string(data)}, " ")
}
