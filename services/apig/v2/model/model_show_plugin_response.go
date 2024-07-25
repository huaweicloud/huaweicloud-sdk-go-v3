package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowPluginResponse Response Object
type ShowPluginResponse struct {

	// 插件编码。
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称。支持汉字，英文，数字，中划线，下划线，且只能以英文和汉字开头，3-255字符。 > 中文字符必须为UTF-8或者unicode编码。
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证 - proxy_cache: 响应缓存
	PluginType *ShowPluginResponsePluginType `json:"plugin_type,omitempty"`

	// 插件可见范围。global：全局可见；
	PluginScope *ShowPluginResponsePluginScope `json:"plugin_scope,omitempty"`

	// 插件定义内容，支持json。参考提供的具体模型定义  CorsPluginContent：跨域资源共享 定义内容 SetRespHeadersContent：HTTP响应头管理 定义内容 KafkaLogContent：Kafka日志推送 定义内容 BreakerContent：断路器 定义内容 RateLimitContent 流量控制 定义内容 ThirdAuthContent: 第三方认证 定义内容 ProxyCacheContent: 响应缓存 定义内容
	PluginContent *string `json:"plugin_content,omitempty"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginResponse", string(data)}, " ")
}

type ShowPluginResponsePluginType struct {
	value string
}

type ShowPluginResponsePluginTypeEnum struct {
	CORS             ShowPluginResponsePluginType
	SET_RESP_HEADERS ShowPluginResponsePluginType
	KAFKA_LOG        ShowPluginResponsePluginType
	BREAKER          ShowPluginResponsePluginType
	RATE_LIMIT       ShowPluginResponsePluginType
	THIRD_AUTH       ShowPluginResponsePluginType
	PROXY_CACHE      ShowPluginResponsePluginType
}

func GetShowPluginResponsePluginTypeEnum() ShowPluginResponsePluginTypeEnum {
	return ShowPluginResponsePluginTypeEnum{
		CORS: ShowPluginResponsePluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: ShowPluginResponsePluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: ShowPluginResponsePluginType{
			value: "kafka_log",
		},
		BREAKER: ShowPluginResponsePluginType{
			value: "breaker",
		},
		RATE_LIMIT: ShowPluginResponsePluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: ShowPluginResponsePluginType{
			value: "third_auth",
		},
		PROXY_CACHE: ShowPluginResponsePluginType{
			value: "proxy_cache",
		},
	}
}

func (c ShowPluginResponsePluginType) Value() string {
	return c.value
}

func (c ShowPluginResponsePluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPluginResponsePluginType) UnmarshalJSON(b []byte) error {
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

type ShowPluginResponsePluginScope struct {
	value string
}

type ShowPluginResponsePluginScopeEnum struct {
	GLOBAL ShowPluginResponsePluginScope
}

func GetShowPluginResponsePluginScopeEnum() ShowPluginResponsePluginScopeEnum {
	return ShowPluginResponsePluginScopeEnum{
		GLOBAL: ShowPluginResponsePluginScope{
			value: "global",
		},
	}
}

func (c ShowPluginResponsePluginScope) Value() string {
	return c.value
}

func (c ShowPluginResponsePluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPluginResponsePluginScope) UnmarshalJSON(b []byte) error {
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
