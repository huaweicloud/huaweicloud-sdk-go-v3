package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type AttachedPluginInfo struct {

	// 插件绑定编码。
	PluginAttachId *string `json:"plugin_attach_id,omitempty"`

	// 插件编码。
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称。支持汉字，英文，数字，中划线，下划线，且只能以英文和汉字开头，3-255字符 > 中文字符必须为UTF-8或者unicode编码。
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证 - proxy_cache: 响应缓存
	PluginType *AttachedPluginInfoPluginType `json:"plugin_type,omitempty"`

	// 插件可见范围。global：全局可见。
	PluginScope *AttachedPluginInfoPluginScope `json:"plugin_scope,omitempty"`

	// 绑定API的环境编码。
	EnvId *string `json:"env_id,omitempty"`

	// api授权绑定的环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 绑定时间。
	AttachedTime *sdktime.SdkTime `json:"attached_time,omitempty"`

	// 插件定义内容，支持json。
	PluginContent *string `json:"plugin_content,omitempty"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o AttachedPluginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedPluginInfo struct{}"
	}

	return strings.Join([]string{"AttachedPluginInfo", string(data)}, " ")
}

type AttachedPluginInfoPluginType struct {
	value string
}

type AttachedPluginInfoPluginTypeEnum struct {
	CORS             AttachedPluginInfoPluginType
	SET_RESP_HEADERS AttachedPluginInfoPluginType
	KAFKA_LOG        AttachedPluginInfoPluginType
	BREAKER          AttachedPluginInfoPluginType
	RATE_LIMIT       AttachedPluginInfoPluginType
	THIRD_AUTH       AttachedPluginInfoPluginType
	PROXY_CACHE      AttachedPluginInfoPluginType
}

func GetAttachedPluginInfoPluginTypeEnum() AttachedPluginInfoPluginTypeEnum {
	return AttachedPluginInfoPluginTypeEnum{
		CORS: AttachedPluginInfoPluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: AttachedPluginInfoPluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: AttachedPluginInfoPluginType{
			value: "kafka_log",
		},
		BREAKER: AttachedPluginInfoPluginType{
			value: "breaker",
		},
		RATE_LIMIT: AttachedPluginInfoPluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: AttachedPluginInfoPluginType{
			value: "third_auth",
		},
		PROXY_CACHE: AttachedPluginInfoPluginType{
			value: "proxy_cache",
		},
	}
}

func (c AttachedPluginInfoPluginType) Value() string {
	return c.value
}

func (c AttachedPluginInfoPluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachedPluginInfoPluginType) UnmarshalJSON(b []byte) error {
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

type AttachedPluginInfoPluginScope struct {
	value string
}

type AttachedPluginInfoPluginScopeEnum struct {
	GLOBAL AttachedPluginInfoPluginScope
}

func GetAttachedPluginInfoPluginScopeEnum() AttachedPluginInfoPluginScopeEnum {
	return AttachedPluginInfoPluginScopeEnum{
		GLOBAL: AttachedPluginInfoPluginScope{
			value: "global",
		},
	}
}

func (c AttachedPluginInfoPluginScope) Value() string {
	return c.value
}

func (c AttachedPluginInfoPluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachedPluginInfoPluginScope) UnmarshalJSON(b []byte) error {
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
