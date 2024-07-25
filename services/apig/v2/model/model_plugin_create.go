package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PluginCreate struct {

	// 插件名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3-255字符。 > 中文字符必须为UTF-8或者unicode编码。
	PluginName string `json:"plugin_name"`

	// 插件类型 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送  - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证 - proxy_cache: 响应缓存
	PluginType PluginCreatePluginType `json:"plugin_type"`

	// 插件可见范围。global：全局可见；
	PluginScope PluginCreatePluginScope `json:"plugin_scope"`

	// 插件定义内容，支持json。参考提供的具体模型定义  CorsPluginContent：跨域资源共享 定义内容 SetRespHeadersContent：HTTP响应头管理 定义内容 KafkaLogContent：Kafka日志推送 定义内容 BreakerContent：断路器 定义内容 RateLimitContent 流量控制 定义内容 ThirdAuthContent: 第三方认证 定义内容 ProxyCacheContent: 响应缓存 定义内容
	PluginContent string `json:"plugin_content"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`
}

func (o PluginCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginCreate struct{}"
	}

	return strings.Join([]string{"PluginCreate", string(data)}, " ")
}

type PluginCreatePluginType struct {
	value string
}

type PluginCreatePluginTypeEnum struct {
	CORS             PluginCreatePluginType
	SET_RESP_HEADERS PluginCreatePluginType
	KAFKA_LOG        PluginCreatePluginType
	BREAKER          PluginCreatePluginType
	RATE_LIMIT       PluginCreatePluginType
	THIRD_AUTH       PluginCreatePluginType
	PROXY_CACHE      PluginCreatePluginType
}

func GetPluginCreatePluginTypeEnum() PluginCreatePluginTypeEnum {
	return PluginCreatePluginTypeEnum{
		CORS: PluginCreatePluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: PluginCreatePluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: PluginCreatePluginType{
			value: "kafka_log",
		},
		BREAKER: PluginCreatePluginType{
			value: "breaker",
		},
		RATE_LIMIT: PluginCreatePluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: PluginCreatePluginType{
			value: "third_auth",
		},
		PROXY_CACHE: PluginCreatePluginType{
			value: "proxy_cache",
		},
	}
}

func (c PluginCreatePluginType) Value() string {
	return c.value
}

func (c PluginCreatePluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginCreatePluginType) UnmarshalJSON(b []byte) error {
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

type PluginCreatePluginScope struct {
	value string
}

type PluginCreatePluginScopeEnum struct {
	GLOBAL PluginCreatePluginScope
}

func GetPluginCreatePluginScopeEnum() PluginCreatePluginScopeEnum {
	return PluginCreatePluginScopeEnum{
		GLOBAL: PluginCreatePluginScope{
			value: "global",
		},
	}
}

func (c PluginCreatePluginScope) Value() string {
	return c.value
}

func (c PluginCreatePluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginCreatePluginScope) UnmarshalJSON(b []byte) error {
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
