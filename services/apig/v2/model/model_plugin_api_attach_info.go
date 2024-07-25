package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type PluginApiAttachInfo struct {

	// 插件绑定编码。
	PluginAttachId *string `json:"plugin_attach_id,omitempty"`

	// 插件编码。
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称。支持汉字，英文，数字，中划线，下划线，且只能以英文和汉字开头，3-255字符 > 中文字符必须为UTF-8或者unicode编码。
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证 - proxy_cache: 响应缓存
	PluginType *PluginApiAttachInfoPluginType `json:"plugin_type,omitempty"`

	// 插件可见范围。global：全局可见。
	PluginScope *PluginApiAttachInfoPluginScope `json:"plugin_scope,omitempty"`

	// 绑定API的环境编码。
	EnvId *string `json:"env_id,omitempty"`

	// api授权绑定的环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 绑定的API编码。
	ApiId *string `json:"api_id,omitempty"`

	// API的名称
	ApiName *string `json:"api_name,omitempty"`

	// 绑定时间。
	AttachedTime *sdktime.SdkTime `json:"attached_time,omitempty"`
}

func (o PluginApiAttachInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginApiAttachInfo struct{}"
	}

	return strings.Join([]string{"PluginApiAttachInfo", string(data)}, " ")
}

type PluginApiAttachInfoPluginType struct {
	value string
}

type PluginApiAttachInfoPluginTypeEnum struct {
	CORS             PluginApiAttachInfoPluginType
	SET_RESP_HEADERS PluginApiAttachInfoPluginType
	KAFKA_LOG        PluginApiAttachInfoPluginType
	BREAKER          PluginApiAttachInfoPluginType
	RATE_LIMIT       PluginApiAttachInfoPluginType
	THIRD_AUTH       PluginApiAttachInfoPluginType
	PROXY_CACHE      PluginApiAttachInfoPluginType
}

func GetPluginApiAttachInfoPluginTypeEnum() PluginApiAttachInfoPluginTypeEnum {
	return PluginApiAttachInfoPluginTypeEnum{
		CORS: PluginApiAttachInfoPluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: PluginApiAttachInfoPluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: PluginApiAttachInfoPluginType{
			value: "kafka_log",
		},
		BREAKER: PluginApiAttachInfoPluginType{
			value: "breaker",
		},
		RATE_LIMIT: PluginApiAttachInfoPluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: PluginApiAttachInfoPluginType{
			value: "third_auth",
		},
		PROXY_CACHE: PluginApiAttachInfoPluginType{
			value: "proxy_cache",
		},
	}
}

func (c PluginApiAttachInfoPluginType) Value() string {
	return c.value
}

func (c PluginApiAttachInfoPluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginApiAttachInfoPluginType) UnmarshalJSON(b []byte) error {
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

type PluginApiAttachInfoPluginScope struct {
	value string
}

type PluginApiAttachInfoPluginScopeEnum struct {
	GLOBAL PluginApiAttachInfoPluginScope
}

func GetPluginApiAttachInfoPluginScopeEnum() PluginApiAttachInfoPluginScopeEnum {
	return PluginApiAttachInfoPluginScopeEnum{
		GLOBAL: PluginApiAttachInfoPluginScope{
			value: "global",
		},
	}
}

func (c PluginApiAttachInfoPluginScope) Value() string {
	return c.value
}

func (c PluginApiAttachInfoPluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginApiAttachInfoPluginScope) UnmarshalJSON(b []byte) error {
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
