package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowDetailsOfCustomAuthorizersV2Response struct {

	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。
	Name string `json:"name" xml:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端  不支持修改
	Type ShowDetailsOfCustomAuthorizersV2ResponseType `json:"type" xml:"type"`

	// 只能为：FUNC
	AuthorizerType ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType `json:"authorizer_type" xml:"authorizer_type"`

	// 函数地址。
	AuthorizerUri string `json:"authorizer_uri" xml:"authorizer_uri"`

	// 认证来源
	Identities *[]Identity `json:"identities,omitempty" xml:"identities"`

	// 缓存时间
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 用户数据
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 自定义后端服务ID。  暂不支持
	LdApiId *string `json:"ld_api_id,omitempty" xml:"ld_api_id"`

	// 是否发送body
	NeedBody *bool `json:"need_body,omitempty" xml:"need_body"`

	// 自定义认证编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 自定义认证所属应用编号  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty" xml:"roma_app_id"`

	// 自定义认证所属应用名称  暂不支持
	RomaAppName    *string `json:"roma_app_name,omitempty" xml:"roma_app_name"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailsOfCustomAuthorizersV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfCustomAuthorizersV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfCustomAuthorizersV2Response", string(data)}, " ")
}

type ShowDetailsOfCustomAuthorizersV2ResponseType struct {
	value string
}

type ShowDetailsOfCustomAuthorizersV2ResponseTypeEnum struct {
	FRONTEND ShowDetailsOfCustomAuthorizersV2ResponseType
	BACKEND  ShowDetailsOfCustomAuthorizersV2ResponseType
}

func GetShowDetailsOfCustomAuthorizersV2ResponseTypeEnum() ShowDetailsOfCustomAuthorizersV2ResponseTypeEnum {
	return ShowDetailsOfCustomAuthorizersV2ResponseTypeEnum{
		FRONTEND: ShowDetailsOfCustomAuthorizersV2ResponseType{
			value: "FRONTEND",
		},
		BACKEND: ShowDetailsOfCustomAuthorizersV2ResponseType{
			value: "BACKEND",
		},
	}
}

func (c ShowDetailsOfCustomAuthorizersV2ResponseType) Value() string {
	return c.value
}

func (c ShowDetailsOfCustomAuthorizersV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfCustomAuthorizersV2ResponseType) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType struct {
	value string
}

type ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerTypeEnum struct {
	FUNC ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType
}

func GetShowDetailsOfCustomAuthorizersV2ResponseAuthorizerTypeEnum() ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerTypeEnum {
	return ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerTypeEnum{
		FUNC: ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType) Value() string {
	return c.value
}

func (c ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfCustomAuthorizersV2ResponseAuthorizerType) UnmarshalJSON(b []byte) error {
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
