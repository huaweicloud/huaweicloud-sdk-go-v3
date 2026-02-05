package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MemberGroupUrlInfo struct {

	// 后端服务器分组ID。
	MemberGroupId string `json:"member_group_id"`

	// 请求协议。
	ReqProtocol MemberGroupUrlInfoReqProtocol `json:"req_protocol"`

	// 请求方式。
	ReqMethod MemberGroupUrlInfoReqMethod `json:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512字符，且满足URI规范。 支持环境变量，使用环境变量时，每个变量名的长度为3~32位的字符串，字符串由英文字母、数字、中划线、下划线组成，且只能以英文开头。 > 需要服从URI规范。
	ReqUri string `json:"req_uri"`
}

func (o MemberGroupUrlInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberGroupUrlInfo struct{}"
	}

	return strings.Join([]string{"MemberGroupUrlInfo", string(data)}, " ")
}

type MemberGroupUrlInfoReqProtocol struct {
	value string
}

type MemberGroupUrlInfoReqProtocolEnum struct {
	HTTP  MemberGroupUrlInfoReqProtocol
	HTTPS MemberGroupUrlInfoReqProtocol
	GRPC  MemberGroupUrlInfoReqProtocol
	GRPCS MemberGroupUrlInfoReqProtocol
}

func GetMemberGroupUrlInfoReqProtocolEnum() MemberGroupUrlInfoReqProtocolEnum {
	return MemberGroupUrlInfoReqProtocolEnum{
		HTTP: MemberGroupUrlInfoReqProtocol{
			value: "HTTP",
		},
		HTTPS: MemberGroupUrlInfoReqProtocol{
			value: "HTTPS",
		},
		GRPC: MemberGroupUrlInfoReqProtocol{
			value: "GRPC",
		},
		GRPCS: MemberGroupUrlInfoReqProtocol{
			value: "GRPCS",
		},
	}
}

func (c MemberGroupUrlInfoReqProtocol) Value() string {
	return c.value
}

func (c MemberGroupUrlInfoReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberGroupUrlInfoReqProtocol) UnmarshalJSON(b []byte) error {
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

type MemberGroupUrlInfoReqMethod struct {
	value string
}

type MemberGroupUrlInfoReqMethodEnum struct {
	GET     MemberGroupUrlInfoReqMethod
	POST    MemberGroupUrlInfoReqMethod
	PUT     MemberGroupUrlInfoReqMethod
	DELETE  MemberGroupUrlInfoReqMethod
	HEAD    MemberGroupUrlInfoReqMethod
	PATCH   MemberGroupUrlInfoReqMethod
	OPTIONS MemberGroupUrlInfoReqMethod
	ANY     MemberGroupUrlInfoReqMethod
}

func GetMemberGroupUrlInfoReqMethodEnum() MemberGroupUrlInfoReqMethodEnum {
	return MemberGroupUrlInfoReqMethodEnum{
		GET: MemberGroupUrlInfoReqMethod{
			value: "GET",
		},
		POST: MemberGroupUrlInfoReqMethod{
			value: "POST",
		},
		PUT: MemberGroupUrlInfoReqMethod{
			value: "PUT",
		},
		DELETE: MemberGroupUrlInfoReqMethod{
			value: "DELETE",
		},
		HEAD: MemberGroupUrlInfoReqMethod{
			value: "HEAD",
		},
		PATCH: MemberGroupUrlInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: MemberGroupUrlInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: MemberGroupUrlInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c MemberGroupUrlInfoReqMethod) Value() string {
	return c.value
}

func (c MemberGroupUrlInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberGroupUrlInfoReqMethod) UnmarshalJSON(b []byte) error {
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
