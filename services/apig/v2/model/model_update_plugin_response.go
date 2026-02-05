package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdatePluginResponse Response Object
type UpdatePluginResponse struct {

	// 插件编码。
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称。支持汉字，英文，数字，中划线，下划线，且只能以英文和汉字开头，3-255字符。 > 中文字符必须为UTF-8或者unicode编码。
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型。 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证 - proxy_cache: 响应缓存 - proxy_mirror: 请求镜像 - oidc_auth: OIDC认证 - jwt_auth: JWT认证
	PluginType *UpdatePluginResponsePluginType `json:"plugin_type,omitempty"`

	// 插件可见范围。global：全局可见；
	PluginScope *UpdatePluginResponsePluginScope `json:"plugin_scope,omitempty"`

	// 插件定义内容，支持json。参考提供的具体模型定义  [CorsPluginContent](apig-api-CorsPluginContent.xml)：跨域资源共享 定义内容 [SetRespHeadersContent](apig-api-SetRespHeadersContent.xml)：HTTP响应头管理 定义内容 [KafkaLogContent](apig-api-KafkaLogContent.xml)：Kafka日志推送 定义内容 [BreakerContent](apig-api-BreakerContent.xml)：断路器 定义内容 [RateLimitContent](apig-api-RateLimitContent.xml)：流量控制 定义内容 [ThirdAuthContent](apig-api-ThirdAuthContent.xml)：第三方认证 定义内容 [ProxyCacheContent](apig-api-ProxyCacheContent.xml)：响应缓存 定义内容 [ProxyMirrorContent](apig-api-ProxyMirrorContent.xml)：请求镜像 定义内容 [OIDCAuthResp](apig-api-OIDCAuthContent.xml)：OIDC认证 定义内容 [JWTAuthContent](apig-api-JWTAuthContent.xml)：JWT认证 定义内容
	PluginContent *string `json:"plugin_content,omitempty"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdatePluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginResponse struct{}"
	}

	return strings.Join([]string{"UpdatePluginResponse", string(data)}, " ")
}

type UpdatePluginResponsePluginType struct {
	value string
}

type UpdatePluginResponsePluginTypeEnum struct {
	CORS             UpdatePluginResponsePluginType
	SET_RESP_HEADERS UpdatePluginResponsePluginType
	KAFKA_LOG        UpdatePluginResponsePluginType
	BREAKER          UpdatePluginResponsePluginType
	RATE_LIMIT       UpdatePluginResponsePluginType
	THIRD_AUTH       UpdatePluginResponsePluginType
	PROXY_CACHE      UpdatePluginResponsePluginType
	PROXY_MIRROR     UpdatePluginResponsePluginType
	OIDC_AUTH        UpdatePluginResponsePluginType
	JWT_AUTH         UpdatePluginResponsePluginType
}

func GetUpdatePluginResponsePluginTypeEnum() UpdatePluginResponsePluginTypeEnum {
	return UpdatePluginResponsePluginTypeEnum{
		CORS: UpdatePluginResponsePluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: UpdatePluginResponsePluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: UpdatePluginResponsePluginType{
			value: "kafka_log",
		},
		BREAKER: UpdatePluginResponsePluginType{
			value: "breaker",
		},
		RATE_LIMIT: UpdatePluginResponsePluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: UpdatePluginResponsePluginType{
			value: "third_auth",
		},
		PROXY_CACHE: UpdatePluginResponsePluginType{
			value: "proxy_cache",
		},
		PROXY_MIRROR: UpdatePluginResponsePluginType{
			value: "proxy_mirror",
		},
		OIDC_AUTH: UpdatePluginResponsePluginType{
			value: "oidc_auth",
		},
		JWT_AUTH: UpdatePluginResponsePluginType{
			value: "jwt_auth",
		},
	}
}

func (c UpdatePluginResponsePluginType) Value() string {
	return c.value
}

func (c UpdatePluginResponsePluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePluginResponsePluginType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdatePluginResponsePluginScope struct {
	value string
}

type UpdatePluginResponsePluginScopeEnum struct {
	GLOBAL UpdatePluginResponsePluginScope
}

func GetUpdatePluginResponsePluginScopeEnum() UpdatePluginResponsePluginScopeEnum {
	return UpdatePluginResponsePluginScopeEnum{
		GLOBAL: UpdatePluginResponsePluginScope{
			value: "global",
		},
	}
}

func (c UpdatePluginResponsePluginScope) Value() string {
	return c.value
}

func (c UpdatePluginResponsePluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePluginResponsePluginScope) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
