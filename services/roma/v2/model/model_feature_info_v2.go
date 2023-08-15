package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type FeatureInfoV2 struct {

	// 特性编号
	Id *string `json:"id,omitempty"`

	// 特性名称
	Name *FeatureInfoV2Name `json:"name,omitempty"`

	// 是否开启特性
	Enable *bool `json:"enable,omitempty"`

	// 特性参数配置
	Config *string `json:"config,omitempty"`

	// 实例编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例特性更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o FeatureInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureInfoV2 struct{}"
	}

	return strings.Join([]string{"FeatureInfoV2", string(data)}, " ")
}

type FeatureInfoV2Name struct {
	value string
}

type FeatureInfoV2NameEnum struct {
	LTS                         FeatureInfoV2Name
	GATEWAY_RESPONSES           FeatureInfoV2Name
	RATELIMIT                   FeatureInfoV2Name
	REQUEST_BODY_SIZE           FeatureInfoV2Name
	BACKEND_TIMEOUT             FeatureInfoV2Name
	APP_TOKEN                   FeatureInfoV2Name
	APP_API_KEY                 FeatureInfoV2Name
	APP_BASIC                   FeatureInfoV2Name
	APP_SECRET                  FeatureInfoV2Name
	APP_JWT                     FeatureInfoV2Name
	PUBLIC_KEY                  FeatureInfoV2Name
	BACKEND_TOKEN_ALLOW         FeatureInfoV2Name
	SIGN_BASIC                  FeatureInfoV2Name
	MULTI_AUTH                  FeatureInfoV2Name
	BACKEND_CLIENT_CERTIFICATE  FeatureInfoV2Name
	SSL_CIPHERS                 FeatureInfoV2Name
	APP_CONFIG                  FeatureInfoV2Name
	GREEN_TUNNEL                FeatureInfoV2Name
	APP_ROUTE                   FeatureInfoV2Name
	DEFAULT_GROUP_HIDE          FeatureInfoV2Name
	CASCADE                     FeatureInfoV2Name
	SANDBOX                     FeatureInfoV2Name
	LIVEDATA_CONFIG             FeatureInfoV2Name
	APICLIENT_FIRST_USE_X_HW_ID FeatureInfoV2Name
	CORS                        FeatureInfoV2Name
	API_TASK                    FeatureInfoV2Name
	APP_QUOTA                   FeatureInfoV2Name
	APP_ACL                     FeatureInfoV2Name
	VPC_BACKUP                  FeatureInfoV2Name
	THROTTLE_STRATEGY           FeatureInfoV2Name
}

func GetFeatureInfoV2NameEnum() FeatureInfoV2NameEnum {
	return FeatureInfoV2NameEnum{
		LTS: FeatureInfoV2Name{
			value: "lts",
		},
		GATEWAY_RESPONSES: FeatureInfoV2Name{
			value: "gateway_responses",
		},
		RATELIMIT: FeatureInfoV2Name{
			value: "ratelimit",
		},
		REQUEST_BODY_SIZE: FeatureInfoV2Name{
			value: "request_body_size",
		},
		BACKEND_TIMEOUT: FeatureInfoV2Name{
			value: "backend_timeout",
		},
		APP_TOKEN: FeatureInfoV2Name{
			value: "app_token",
		},
		APP_API_KEY: FeatureInfoV2Name{
			value: "app_api_key",
		},
		APP_BASIC: FeatureInfoV2Name{
			value: "app_basic",
		},
		APP_SECRET: FeatureInfoV2Name{
			value: "app_secret",
		},
		APP_JWT: FeatureInfoV2Name{
			value: "app_jwt",
		},
		PUBLIC_KEY: FeatureInfoV2Name{
			value: "public_key",
		},
		BACKEND_TOKEN_ALLOW: FeatureInfoV2Name{
			value: "backend_token_allow",
		},
		SIGN_BASIC: FeatureInfoV2Name{
			value: "sign_basic",
		},
		MULTI_AUTH: FeatureInfoV2Name{
			value: "multi_auth",
		},
		BACKEND_CLIENT_CERTIFICATE: FeatureInfoV2Name{
			value: "backend_client_certificate",
		},
		SSL_CIPHERS: FeatureInfoV2Name{
			value: "ssl_ciphers",
		},
		APP_CONFIG: FeatureInfoV2Name{
			value: "app_config",
		},
		GREEN_TUNNEL: FeatureInfoV2Name{
			value: "green_tunnel",
		},
		APP_ROUTE: FeatureInfoV2Name{
			value: "app_route",
		},
		DEFAULT_GROUP_HIDE: FeatureInfoV2Name{
			value: "default_group_hide",
		},
		CASCADE: FeatureInfoV2Name{
			value: "cascade",
		},
		SANDBOX: FeatureInfoV2Name{
			value: "sandbox",
		},
		LIVEDATA_CONFIG: FeatureInfoV2Name{
			value: "livedata_config",
		},
		APICLIENT_FIRST_USE_X_HW_ID: FeatureInfoV2Name{
			value: "apiclient_first_use_x_hw_id",
		},
		CORS: FeatureInfoV2Name{
			value: "cors",
		},
		API_TASK: FeatureInfoV2Name{
			value: "api_task",
		},
		APP_QUOTA: FeatureInfoV2Name{
			value: "app_quota",
		},
		APP_ACL: FeatureInfoV2Name{
			value: "app_acl",
		},
		VPC_BACKUP: FeatureInfoV2Name{
			value: "vpc_backup",
		},
		THROTTLE_STRATEGY: FeatureInfoV2Name{
			value: "throttle_strategy",
		},
	}
}

func (c FeatureInfoV2Name) Value() string {
	return c.value
}

func (c FeatureInfoV2Name) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FeatureInfoV2Name) UnmarshalJSON(b []byte) error {
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
