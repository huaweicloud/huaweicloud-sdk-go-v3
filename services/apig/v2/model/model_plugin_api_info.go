package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type PluginApiInfo struct {

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// API名称。   支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。  > 中文字符必须为UTF-8或者unicode编码。
	ApiName *string `json:"api_name,omitempty"`

	// API类型 - 1：公有API - 2：私有API
	Type *int32 `json:"type,omitempty"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS
	ReqProtocol *PluginApiInfoReqProtocol `json:"req_protocol,omitempty"`

	// API的请求方式
	ReqMethod *PluginApiInfoReqMethod `json:"req_method,omitempty"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。   支持环境变量，使用环境变量时，每个变量名的长度为3 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线组成，且只能以英文开头。  > 需要服从URI规范。
	ReqUri *string `json:"req_uri,omitempty"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *PluginApiInfoAuthType `json:"auth_type,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *PluginApiInfoMatchMode `json:"match_mode,omitempty"`

	// API描述。
	Remark *string `json:"remark,omitempty"`

	// API所属的分组编号
	GroupId *string `json:"group_id,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 归属集成应用编码,兼容roma实例的字段，一般为空
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 绑定API的环境编码。
	EnvId *string `json:"env_id,omitempty"`

	// 绑定API的环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 发布编码。
	PublishId *string `json:"publish_id,omitempty"`

	// 插件绑定编码。
	PluginAttachId *string `json:"plugin_attach_id,omitempty"`

	// 绑定时间。
	AttachedTime *sdktime.SdkTime `json:"attached_time,omitempty"`
}

func (o PluginApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginApiInfo struct{}"
	}

	return strings.Join([]string{"PluginApiInfo", string(data)}, " ")
}

type PluginApiInfoReqProtocol struct {
	value string
}

type PluginApiInfoReqProtocolEnum struct {
	HTTP  PluginApiInfoReqProtocol
	HTTPS PluginApiInfoReqProtocol
	BOTH  PluginApiInfoReqProtocol
}

func GetPluginApiInfoReqProtocolEnum() PluginApiInfoReqProtocolEnum {
	return PluginApiInfoReqProtocolEnum{
		HTTP: PluginApiInfoReqProtocol{
			value: "HTTP",
		},
		HTTPS: PluginApiInfoReqProtocol{
			value: "HTTPS",
		},
		BOTH: PluginApiInfoReqProtocol{
			value: "BOTH",
		},
	}
}

func (c PluginApiInfoReqProtocol) Value() string {
	return c.value
}

func (c PluginApiInfoReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginApiInfoReqProtocol) UnmarshalJSON(b []byte) error {
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

type PluginApiInfoReqMethod struct {
	value string
}

type PluginApiInfoReqMethodEnum struct {
	GET     PluginApiInfoReqMethod
	POST    PluginApiInfoReqMethod
	PUT     PluginApiInfoReqMethod
	DELETE  PluginApiInfoReqMethod
	HEAD    PluginApiInfoReqMethod
	PATCH   PluginApiInfoReqMethod
	OPTIONS PluginApiInfoReqMethod
	ANY     PluginApiInfoReqMethod
}

func GetPluginApiInfoReqMethodEnum() PluginApiInfoReqMethodEnum {
	return PluginApiInfoReqMethodEnum{
		GET: PluginApiInfoReqMethod{
			value: "GET",
		},
		POST: PluginApiInfoReqMethod{
			value: "POST",
		},
		PUT: PluginApiInfoReqMethod{
			value: "PUT",
		},
		DELETE: PluginApiInfoReqMethod{
			value: "DELETE",
		},
		HEAD: PluginApiInfoReqMethod{
			value: "HEAD",
		},
		PATCH: PluginApiInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: PluginApiInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: PluginApiInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c PluginApiInfoReqMethod) Value() string {
	return c.value
}

func (c PluginApiInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginApiInfoReqMethod) UnmarshalJSON(b []byte) error {
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

type PluginApiInfoAuthType struct {
	value string
}

type PluginApiInfoAuthTypeEnum struct {
	NONE       PluginApiInfoAuthType
	APP        PluginApiInfoAuthType
	IAM        PluginApiInfoAuthType
	AUTHORIZER PluginApiInfoAuthType
}

func GetPluginApiInfoAuthTypeEnum() PluginApiInfoAuthTypeEnum {
	return PluginApiInfoAuthTypeEnum{
		NONE: PluginApiInfoAuthType{
			value: "NONE",
		},
		APP: PluginApiInfoAuthType{
			value: "APP",
		},
		IAM: PluginApiInfoAuthType{
			value: "IAM",
		},
		AUTHORIZER: PluginApiInfoAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c PluginApiInfoAuthType) Value() string {
	return c.value
}

func (c PluginApiInfoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginApiInfoAuthType) UnmarshalJSON(b []byte) error {
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

type PluginApiInfoMatchMode struct {
	value string
}

type PluginApiInfoMatchModeEnum struct {
	SWA    PluginApiInfoMatchMode
	NORMAL PluginApiInfoMatchMode
}

func GetPluginApiInfoMatchModeEnum() PluginApiInfoMatchModeEnum {
	return PluginApiInfoMatchModeEnum{
		SWA: PluginApiInfoMatchMode{
			value: "SWA",
		},
		NORMAL: PluginApiInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c PluginApiInfoMatchMode) Value() string {
	return c.value
}

func (c PluginApiInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginApiInfoMatchMode) UnmarshalJSON(b []byte) error {
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
