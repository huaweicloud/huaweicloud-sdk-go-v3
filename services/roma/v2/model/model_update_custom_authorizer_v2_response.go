package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateCustomAuthorizerV2Response Response Object
type UpdateCustomAuthorizerV2Response struct {

	// 自定义认证的名称。 长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。
	Name string `json:"name"`

	// 自定义认证类型  - FRONTEND：前端 - BACKEND：后端
	Type UpdateCustomAuthorizerV2ResponseType `json:"type"`

	// 自定义认证函数类型： - LD：自定义后端函数 - FUNC：[函数服务函数](tag:hws,hws_hk,hcs,hcs_sm,fcs,g42)[暂不支持](tag:Site)
	AuthorizerType UpdateCustomAuthorizerV2ResponseAuthorizerType `json:"authorizer_type"`

	// 函数地址。  注意：使用自定义后端的函数API，API请求方法必须为POST，且API状态必须为已部署。
	AuthorizerUri string `json:"authorizer_uri"`

	// 函数版本。  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AuthorizerVersion *string `json:"authorizer_version,omitempty"`

	// 函数别名地址。  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AuthorizerAliasUri *string `json:"authorizer_alias_uri,omitempty"`

	// 认证来源
	Identities *[]Identity `json:"identities,omitempty"`

	// 缓存时间
	Ttl *int32 `json:"ttl,omitempty"`

	// 用户数据
	UserData *string `json:"user_data,omitempty"`

	// 自定义后端服务ID。  自定义认证函数类型为LD时必填
	LdApiId *string `json:"ld_api_id,omitempty"`

	// 是否发送body
	NeedBody *bool `json:"need_body,omitempty"`

	// 自定义认证编号
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 自定义认证所属应用编号
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 自定义认证所属应用名称
	RomaAppName    *string `json:"roma_app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCustomAuthorizerV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomAuthorizerV2Response struct{}"
	}

	return strings.Join([]string{"UpdateCustomAuthorizerV2Response", string(data)}, " ")
}

type UpdateCustomAuthorizerV2ResponseType struct {
	value string
}

type UpdateCustomAuthorizerV2ResponseTypeEnum struct {
	FRONTEND UpdateCustomAuthorizerV2ResponseType
	BACKEND  UpdateCustomAuthorizerV2ResponseType
}

func GetUpdateCustomAuthorizerV2ResponseTypeEnum() UpdateCustomAuthorizerV2ResponseTypeEnum {
	return UpdateCustomAuthorizerV2ResponseTypeEnum{
		FRONTEND: UpdateCustomAuthorizerV2ResponseType{
			value: "FRONTEND",
		},
		BACKEND: UpdateCustomAuthorizerV2ResponseType{
			value: "BACKEND",
		},
	}
}

func (c UpdateCustomAuthorizerV2ResponseType) Value() string {
	return c.value
}

func (c UpdateCustomAuthorizerV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCustomAuthorizerV2ResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateCustomAuthorizerV2ResponseAuthorizerType struct {
	value string
}

type UpdateCustomAuthorizerV2ResponseAuthorizerTypeEnum struct {
	LD   UpdateCustomAuthorizerV2ResponseAuthorizerType
	FUNC UpdateCustomAuthorizerV2ResponseAuthorizerType
}

func GetUpdateCustomAuthorizerV2ResponseAuthorizerTypeEnum() UpdateCustomAuthorizerV2ResponseAuthorizerTypeEnum {
	return UpdateCustomAuthorizerV2ResponseAuthorizerTypeEnum{
		LD: UpdateCustomAuthorizerV2ResponseAuthorizerType{
			value: "LD",
		},
		FUNC: UpdateCustomAuthorizerV2ResponseAuthorizerType{
			value: "FUNC",
		},
	}
}

func (c UpdateCustomAuthorizerV2ResponseAuthorizerType) Value() string {
	return c.value
}

func (c UpdateCustomAuthorizerV2ResponseAuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCustomAuthorizerV2ResponseAuthorizerType) UnmarshalJSON(b []byte) error {
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
