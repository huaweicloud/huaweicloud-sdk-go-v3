package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiCheckInfoV2 struct {
	// API名称。  type = name时必填

	Name *string `json:"name,omitempty"`
	// 请求方式。  type = path时必填

	ReqMethod *ApiCheckInfoV2ReqMethod `json:"req_method,omitempty"`
	// API的访问地址。  type = path时必填

	ReqUri *string `json:"req_uri,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配）  type = path时必填

	MatchMode *ApiCheckInfoV2MatchMode `json:"match_mode,omitempty"`
	// 分组ID。  校验分组下API定义是否重复时必填

	GroupId *string `json:"group_id,omitempty"`
	// 集成应用ID。  校验应用下API定义是否重复时必填

	RomaAppId *string `json:"roma_app_id,omitempty"`
	// 需要对比的API ID

	ApiId *string `json:"api_id,omitempty"`
	// 校验类型：   - path：路径类型   - name：名称类型

	Type *ApiCheckInfoV2Type `json:"type,omitempty"`
}

func (o ApiCheckInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCheckInfoV2 struct{}"
	}

	return strings.Join([]string{"ApiCheckInfoV2", string(data)}, " ")
}

type ApiCheckInfoV2ReqMethod struct {
	value string
}

type ApiCheckInfoV2ReqMethodEnum struct {
	GET     ApiCheckInfoV2ReqMethod
	POST    ApiCheckInfoV2ReqMethod
	PUT     ApiCheckInfoV2ReqMethod
	DELETE  ApiCheckInfoV2ReqMethod
	HEAD    ApiCheckInfoV2ReqMethod
	PATCH   ApiCheckInfoV2ReqMethod
	OPTIONS ApiCheckInfoV2ReqMethod
	ANY     ApiCheckInfoV2ReqMethod
}

func GetApiCheckInfoV2ReqMethodEnum() ApiCheckInfoV2ReqMethodEnum {
	return ApiCheckInfoV2ReqMethodEnum{
		GET: ApiCheckInfoV2ReqMethod{
			value: "GET",
		},
		POST: ApiCheckInfoV2ReqMethod{
			value: "POST",
		},
		PUT: ApiCheckInfoV2ReqMethod{
			value: "PUT",
		},
		DELETE: ApiCheckInfoV2ReqMethod{
			value: "DELETE",
		},
		HEAD: ApiCheckInfoV2ReqMethod{
			value: "HEAD",
		},
		PATCH: ApiCheckInfoV2ReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiCheckInfoV2ReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiCheckInfoV2ReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiCheckInfoV2ReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCheckInfoV2ReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiCheckInfoV2MatchMode struct {
	value string
}

type ApiCheckInfoV2MatchModeEnum struct {
	SWA    ApiCheckInfoV2MatchMode
	NORMAL ApiCheckInfoV2MatchMode
}

func GetApiCheckInfoV2MatchModeEnum() ApiCheckInfoV2MatchModeEnum {
	return ApiCheckInfoV2MatchModeEnum{
		SWA: ApiCheckInfoV2MatchMode{
			value: "SWA",
		},
		NORMAL: ApiCheckInfoV2MatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiCheckInfoV2MatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCheckInfoV2MatchMode) UnmarshalJSON(b []byte) error {
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

type ApiCheckInfoV2Type struct {
	value string
}

type ApiCheckInfoV2TypeEnum struct {
	PATH ApiCheckInfoV2Type
	NAME ApiCheckInfoV2Type
}

func GetApiCheckInfoV2TypeEnum() ApiCheckInfoV2TypeEnum {
	return ApiCheckInfoV2TypeEnum{
		PATH: ApiCheckInfoV2Type{
			value: "path",
		},
		NAME: ApiCheckInfoV2Type{
			value: "name",
		},
	}
}

func (c ApiCheckInfoV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCheckInfoV2Type) UnmarshalJSON(b []byte) error {
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
