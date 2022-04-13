package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiDebugInfo struct {
	// 请求消息体，最长2097152字节

	Body *string `json:"body,omitempty"`
	// 头域参数，每个参数值为字符串数组，每个参数名称有如下约束： - 英文字母、数字、点、中连线组成 - 必须以英文字母开头，最长32字节 - 不支持以\"X-Apig-\"或\"X-Sdk-\"开头，不区分大小写 - 不支持取值为\"X-Stage\"，不区分大小写 - mode为MARKET或CONSUMER时，不支持取值为\"X-Auth-Token\"和\"Authorization\"，不区分大小写 > 头域名称在使用前会被规范化，如：\"x-MY-hEaDer\"会被规范化为\"X-My-Header\"

	Header map[string]string `json:"header,omitempty"`
	// API的请求方法

	Method ApiDebugInfoMethod `json:"method"`
	// 调试模式 - DEVELOPER 调试尚未发布的API定义 - MARKET 调试云市场已购买的API - CONSUMER 调试指定运行环境下的API定义 > DEVELOPER模式，接口调用者必须是API拥有者。    MARKET模式，接口调用者必须是API购买者或拥有者。    CONSUMER模式，接口调用者必须有API在指定环境上的授权信息或是API拥有者。

	Mode string `json:"mode"`
	// API的请求路径，需以\"/\"开头，最大长度1024 > 须符合路径规范，百分号编码格式可被正确解码

	Path string `json:"path"`
	// 查询参数，每个参数值为字符串数组，每个参数名称有如下约束： - 英文字母、数字、点、下划线、中连线组成 - 必须以英文字母开头，最长32字节 - 不支持以\"X-Apig-\"或\"X-Sdk-\"开头，不区分大小写 - 不支持取值为\"X-Stage\"，不区分大小写

	Query map[string]string `json:"query,omitempty"`
	// API的请求协议 - HTTP - HTTPS

	Scheme string `json:"scheme"`
	// 调试请求使用的APP的key

	AppKey *string `json:"app_key,omitempty"`
	// 调试请求使用的APP的密钥

	AppSecret *string `json:"app_secret,omitempty"`
	// API的访问域名，未提供时根据mode的取值使用如下默认值： - DEVELOPER API分组的子域名 - MARKET 云市场为API分组分配的域名 - CONSUMER API分组的子域名

	Domain *string `json:"domain,omitempty"`
	// 调试请求指定的运行环境，仅在mode为CONSUMER时有效，未提供时有如下默认值: - CONSUMER RELEASE

	Stage *string `json:"stage,omitempty"`
}

func (o ApiDebugInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiDebugInfo struct{}"
	}

	return strings.Join([]string{"ApiDebugInfo", string(data)}, " ")
}

type ApiDebugInfoMethod struct {
	value string
}

type ApiDebugInfoMethodEnum struct {
	GET     ApiDebugInfoMethod
	POST    ApiDebugInfoMethod
	PUT     ApiDebugInfoMethod
	DELETE  ApiDebugInfoMethod
	HEAD    ApiDebugInfoMethod
	PATCH   ApiDebugInfoMethod
	OPTIONS ApiDebugInfoMethod
}

func GetApiDebugInfoMethodEnum() ApiDebugInfoMethodEnum {
	return ApiDebugInfoMethodEnum{
		GET: ApiDebugInfoMethod{
			value: "GET",
		},
		POST: ApiDebugInfoMethod{
			value: "POST",
		},
		PUT: ApiDebugInfoMethod{
			value: "PUT",
		},
		DELETE: ApiDebugInfoMethod{
			value: "DELETE",
		},
		HEAD: ApiDebugInfoMethod{
			value: "HEAD",
		},
		PATCH: ApiDebugInfoMethod{
			value: "PATCH",
		},
		OPTIONS: ApiDebugInfoMethod{
			value: "OPTIONS",
		},
	}
}

func (c ApiDebugInfoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiDebugInfoMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
