package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiDuplicationInfo struct {

	// 请求方式
	ReqMethod *ApiDuplicationInfoReqMethod `json:"req_method,omitempty"`

	// API的访问地址
	ReqUri *string `json:"req_uri,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ApiDuplicationInfoMatchMode `json:"match_mode,omitempty"`

	// 该路径下冲突的api列表
	DuplicatedApis *[]DuplicateApiInfo `json:"duplicated_apis,omitempty"`
}

func (o ApiDuplicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiDuplicationInfo struct{}"
	}

	return strings.Join([]string{"ApiDuplicationInfo", string(data)}, " ")
}

type ApiDuplicationInfoReqMethod struct {
	value string
}

type ApiDuplicationInfoReqMethodEnum struct {
	GET     ApiDuplicationInfoReqMethod
	POST    ApiDuplicationInfoReqMethod
	PUT     ApiDuplicationInfoReqMethod
	DELETE  ApiDuplicationInfoReqMethod
	HEAD    ApiDuplicationInfoReqMethod
	PATCH   ApiDuplicationInfoReqMethod
	OPTIONS ApiDuplicationInfoReqMethod
	ANY     ApiDuplicationInfoReqMethod
}

func GetApiDuplicationInfoReqMethodEnum() ApiDuplicationInfoReqMethodEnum {
	return ApiDuplicationInfoReqMethodEnum{
		GET: ApiDuplicationInfoReqMethod{
			value: "GET",
		},
		POST: ApiDuplicationInfoReqMethod{
			value: "POST",
		},
		PUT: ApiDuplicationInfoReqMethod{
			value: "PUT",
		},
		DELETE: ApiDuplicationInfoReqMethod{
			value: "DELETE",
		},
		HEAD: ApiDuplicationInfoReqMethod{
			value: "HEAD",
		},
		PATCH: ApiDuplicationInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiDuplicationInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiDuplicationInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiDuplicationInfoReqMethod) Value() string {
	return c.value
}

func (c ApiDuplicationInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiDuplicationInfoReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiDuplicationInfoMatchMode struct {
	value string
}

type ApiDuplicationInfoMatchModeEnum struct {
	SWA    ApiDuplicationInfoMatchMode
	NORMAL ApiDuplicationInfoMatchMode
}

func GetApiDuplicationInfoMatchModeEnum() ApiDuplicationInfoMatchModeEnum {
	return ApiDuplicationInfoMatchModeEnum{
		SWA: ApiDuplicationInfoMatchMode{
			value: "SWA",
		},
		NORMAL: ApiDuplicationInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiDuplicationInfoMatchMode) Value() string {
	return c.value
}

func (c ApiDuplicationInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiDuplicationInfoMatchMode) UnmarshalJSON(b []byte) error {
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
