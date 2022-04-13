package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiCheckInfo struct {
	// API名称。  type = name时必填

	Name *string `json:"name,omitempty"`
	// 请求方式。  type = path时必填

	ReqMethod *ApiCheckInfoReqMethod `json:"req_method,omitempty"`
	// API的访问地址。  type = path时必填

	ReqUri *string `json:"req_uri,omitempty"`
	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配）  type = path时必填

	MatchMode *ApiCheckInfoMatchMode `json:"match_mode,omitempty"`
	// 分组ID。  校验分组下API定义是否重复时必填

	GroupId *string `json:"group_id,omitempty"`
	// 集成应用ID。  校验应用下API定义是否重复时必填

	RomaAppId *string `json:"roma_app_id,omitempty"`
	// 需要对比的API ID

	ApiId *string `json:"api_id,omitempty"`
}

func (o ApiCheckInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCheckInfo struct{}"
	}

	return strings.Join([]string{"ApiCheckInfo", string(data)}, " ")
}

type ApiCheckInfoReqMethod struct {
	value string
}

type ApiCheckInfoReqMethodEnum struct {
	GET     ApiCheckInfoReqMethod
	POST    ApiCheckInfoReqMethod
	PUT     ApiCheckInfoReqMethod
	DELETE  ApiCheckInfoReqMethod
	HEAD    ApiCheckInfoReqMethod
	PATCH   ApiCheckInfoReqMethod
	OPTIONS ApiCheckInfoReqMethod
	ANY     ApiCheckInfoReqMethod
}

func GetApiCheckInfoReqMethodEnum() ApiCheckInfoReqMethodEnum {
	return ApiCheckInfoReqMethodEnum{
		GET: ApiCheckInfoReqMethod{
			value: "GET",
		},
		POST: ApiCheckInfoReqMethod{
			value: "POST",
		},
		PUT: ApiCheckInfoReqMethod{
			value: "PUT",
		},
		DELETE: ApiCheckInfoReqMethod{
			value: "DELETE",
		},
		HEAD: ApiCheckInfoReqMethod{
			value: "HEAD",
		},
		PATCH: ApiCheckInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiCheckInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiCheckInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiCheckInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCheckInfoReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiCheckInfoMatchMode struct {
	value string
}

type ApiCheckInfoMatchModeEnum struct {
	SWA    ApiCheckInfoMatchMode
	NORMAL ApiCheckInfoMatchMode
}

func GetApiCheckInfoMatchModeEnum() ApiCheckInfoMatchModeEnum {
	return ApiCheckInfoMatchModeEnum{
		SWA: ApiCheckInfoMatchMode{
			value: "SWA",
		},
		NORMAL: ApiCheckInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiCheckInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCheckInfoMatchMode) UnmarshalJSON(b []byte) error {
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
