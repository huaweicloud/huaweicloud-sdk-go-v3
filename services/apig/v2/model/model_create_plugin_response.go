package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreatePluginResponse Response Object
type CreatePluginResponse struct {

	// 插件编码。
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称。支持汉字，英文，数字，中划线，下划线，且只能以英文和汉字开头，3-255字符。 > 中文字符必须为UTF-8或者unicode编码。
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型。 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证 - proxy_cache: 响应缓存 - proxy_mirror: 请求镜像 - oidc_auth: OIDC认证 - jwt_auth: JWT认证
	PluginType *CreatePluginResponsePluginType `json:"plugin_type,omitempty"`

	// 插件可见范围。global：全局可见；
	PluginScope *CreatePluginResponsePluginScope `json:"plugin_scope,omitempty"`

	// 插件定义内容，支持json。参考提供的具体模型定义  [CorsPluginContent](apig-api-CorsPluginContent.xml)：跨域资源共享 定义内容 [SetRespHeadersContent](apig-api-SetRespHeadersContent.xml)：HTTP响应头管理 定义内容 [KafkaLogContent](apig-api-KafkaLogContent.xml)：Kafka日志推送 定义内容 [BreakerContent](apig-api-BreakerContent.xml)：断路器 定义内容 [RateLimitContent](apig-api-RateLimitContent.xml)：流量控制 定义内容 [ThirdAuthContent](apig-api-ThirdAuthContent.xml)：第三方认证 定义内容 [ProxyCacheContent](apig-api-ProxyCacheContent.xml)：响应缓存 定义内容 [ProxyMirrorContent](apig-api-ProxyMirrorContent.xml)：请求镜像 定义内容 [OIDCAuthContent](apig-api-OIDCAuthContent.xml)：OIDC认证 定义内容 [JWTAuthContent](apig-api-JWTAuthContent.xml)：JWT认证 定义内容
	PluginContent *string `json:"plugin_content,omitempty"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginResponse struct{}"
	}

	return strings.Join([]string{"CreatePluginResponse", string(data)}, " ")
}

type CreatePluginResponsePluginType struct {
	value string
}

type CreatePluginResponsePluginTypeEnum struct {
	CORS             CreatePluginResponsePluginType
	SET_RESP_HEADERS CreatePluginResponsePluginType
	KAFKA_LOG        CreatePluginResponsePluginType
	BREAKER          CreatePluginResponsePluginType
	RATE_LIMIT       CreatePluginResponsePluginType
	THIRD_AUTH       CreatePluginResponsePluginType
	PROXY_CACHE      CreatePluginResponsePluginType
	PROXY_MIRROR     CreatePluginResponsePluginType
	OIDC_AUTH        CreatePluginResponsePluginType
	JWT_AUTH         CreatePluginResponsePluginType
}

func GetCreatePluginResponsePluginTypeEnum() CreatePluginResponsePluginTypeEnum {
	return CreatePluginResponsePluginTypeEnum{
		CORS: CreatePluginResponsePluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: CreatePluginResponsePluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: CreatePluginResponsePluginType{
			value: "kafka_log",
		},
		BREAKER: CreatePluginResponsePluginType{
			value: "breaker",
		},
		RATE_LIMIT: CreatePluginResponsePluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: CreatePluginResponsePluginType{
			value: "third_auth",
		},
		PROXY_CACHE: CreatePluginResponsePluginType{
			value: "proxy_cache",
		},
		PROXY_MIRROR: CreatePluginResponsePluginType{
			value: "proxy_mirror",
		},
		OIDC_AUTH: CreatePluginResponsePluginType{
			value: "oidc_auth",
		},
		JWT_AUTH: CreatePluginResponsePluginType{
			value: "jwt_auth",
		},
	}
}

func (c CreatePluginResponsePluginType) Value() string {
	return c.value
}

func (c CreatePluginResponsePluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePluginResponsePluginType) UnmarshalJSON(b []byte) error {
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

type CreatePluginResponsePluginScope struct {
	value string
}

type CreatePluginResponsePluginScopeEnum struct {
	GLOBAL CreatePluginResponsePluginScope
}

func GetCreatePluginResponsePluginScopeEnum() CreatePluginResponsePluginScopeEnum {
	return CreatePluginResponsePluginScopeEnum{
		GLOBAL: CreatePluginResponsePluginScope{
			value: "global",
		},
	}
}

func (c CreatePluginResponsePluginScope) Value() string {
	return c.value
}

func (c CreatePluginResponsePluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePluginResponsePluginScope) UnmarshalJSON(b []byte) error {
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
