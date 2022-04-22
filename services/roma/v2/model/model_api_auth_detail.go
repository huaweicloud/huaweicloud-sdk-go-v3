package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// api鉴权字段
type ApiAuthDetail struct {

	// 访问API服务的认证方式 - none - basicauth - oauth2.0 - hmac - secret - md5 - apiGateway - keyTop - hikVision - huaweiNetworkManagement - liHe
	AuthMethod *ApiAuthDetailAuthMethod `json:"auth_method,omitempty"`

	// 访问API服务的APP认证方式，认证方式为（apiGateway）时填写 - default - secret - jwt
	AppAuthType *ApiAuthDetailAppAuthType `json:"app_auth_type,omitempty"`

	// 访问API服务的用户名 - 认证方式为（lihe、huaweiNetworkManagement、basicauth）时填写
	UserName *string `json:"user_name,omitempty"`

	// 访问API服务的密码 - 认证方式为（lihe、huaweiNetworkManagement、basicauth、secret、md5、hmac）时填写
	Password *string `json:"password,omitempty"`

	// 访问API服务的AppKey - 认证方式为（apiGateway、oauth2.0）时填写
	AppKey *string `json:"app_key,omitempty"`

	// 访问API服务的AppSecret - 认证方式为（apiGateway、oauth2.0）时填写
	AppSecret *string `json:"app_secret,omitempty"`

	// 访问API服务的Secret - 认证方式为（KeyTop、HikVision、Secret、HMAC、MD5）时填写
	Secret *string `json:"secret,omitempty"`

	// 访问API服务的备用IP - 认证方式为（HuaweiNetworkManagement）时填写
	AltIp *string `json:"alt_ip,omitempty"`

	// 访问API服务的AccessTokenUrl - 认证方式为（liHe、oauth2.0 huaweiNetworkManagement）时填写
	AccessTokenUrl *string `json:"access_token_url,omitempty"`

	// 访问API服务的客户端标识 - 认证方式为Oauth2时填写
	ClientId *string `json:"client_id,omitempty"`

	// 访问API服务的客户端密钥 - 认证方式为Oauth2时填写
	ClientSecret *string `json:"client_secret,omitempty"`

	// 访问API服务的Scope - 认证方式为（LiHe、Oauth2）时填写
	Scope *string `json:"scope,omitempty"`

	// 访问API服务的Authorization - 认证方式为（LiHe）时填写
	Authorization *string `json:"authorization,omitempty"`

	// 访问API服务的授权类型 - 认证方式为（LiHe、Oauth2）时填写 - client_credentials （oauth2.0使用）
	GrantType *string `json:"grant_type,omitempty"`
}

func (o ApiAuthDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthDetail struct{}"
	}

	return strings.Join([]string{"ApiAuthDetail", string(data)}, " ")
}

type ApiAuthDetailAuthMethod struct {
	value string
}

type ApiAuthDetailAuthMethodEnum struct {
	NONE                      ApiAuthDetailAuthMethod
	BASICAUTH                 ApiAuthDetailAuthMethod
	OAUTH2_0                  ApiAuthDetailAuthMethod
	HMAC                      ApiAuthDetailAuthMethod
	SECRET                    ApiAuthDetailAuthMethod
	MD5                       ApiAuthDetailAuthMethod
	API_GATEWAY               ApiAuthDetailAuthMethod
	KEY_TOP                   ApiAuthDetailAuthMethod
	HIK_VISION                ApiAuthDetailAuthMethod
	HUAWEI_NETWORK_MANAGEMENT ApiAuthDetailAuthMethod
	LI_HE                     ApiAuthDetailAuthMethod
}

func GetApiAuthDetailAuthMethodEnum() ApiAuthDetailAuthMethodEnum {
	return ApiAuthDetailAuthMethodEnum{
		NONE: ApiAuthDetailAuthMethod{
			value: "none",
		},
		BASICAUTH: ApiAuthDetailAuthMethod{
			value: "basicauth",
		},
		OAUTH2_0: ApiAuthDetailAuthMethod{
			value: "oauth2.0",
		},
		HMAC: ApiAuthDetailAuthMethod{
			value: "hmac",
		},
		SECRET: ApiAuthDetailAuthMethod{
			value: "secret",
		},
		MD5: ApiAuthDetailAuthMethod{
			value: "md5",
		},
		API_GATEWAY: ApiAuthDetailAuthMethod{
			value: "apiGateway",
		},
		KEY_TOP: ApiAuthDetailAuthMethod{
			value: "keyTop",
		},
		HIK_VISION: ApiAuthDetailAuthMethod{
			value: "hikVision",
		},
		HUAWEI_NETWORK_MANAGEMENT: ApiAuthDetailAuthMethod{
			value: "huaweiNetworkManagement",
		},
		LI_HE: ApiAuthDetailAuthMethod{
			value: "liHe",
		},
	}
}

func (c ApiAuthDetailAuthMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthDetailAuthMethod) UnmarshalJSON(b []byte) error {
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

type ApiAuthDetailAppAuthType struct {
	value string
}

type ApiAuthDetailAppAuthTypeEnum struct {
	DEFAULT ApiAuthDetailAppAuthType
	SECRET  ApiAuthDetailAppAuthType
	JWT     ApiAuthDetailAppAuthType
}

func GetApiAuthDetailAppAuthTypeEnum() ApiAuthDetailAppAuthTypeEnum {
	return ApiAuthDetailAppAuthTypeEnum{
		DEFAULT: ApiAuthDetailAppAuthType{
			value: "default",
		},
		SECRET: ApiAuthDetailAppAuthType{
			value: "secret",
		},
		JWT: ApiAuthDetailAppAuthType{
			value: "jwt",
		},
	}
}

func (c ApiAuthDetailAppAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthDetailAppAuthType) UnmarshalJSON(b []byte) error {
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
