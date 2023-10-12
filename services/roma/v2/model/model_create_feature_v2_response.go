package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateFeatureV2Response Response Object
type CreateFeatureV2Response struct {

	// 特性编号
	Id *string `json:"id,omitempty"`

	// 特性名称
	Name *CreateFeatureV2ResponseName `json:"name,omitempty"`

	// 是否开启特性
	Enable *bool `json:"enable,omitempty"`

	// 特性参数配置
	Config *string `json:"config,omitempty"`

	// 实例编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例特性更新时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateFeatureV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeatureV2Response struct{}"
	}

	return strings.Join([]string{"CreateFeatureV2Response", string(data)}, " ")
}

type CreateFeatureV2ResponseName struct {
	value string
}

type CreateFeatureV2ResponseNameEnum struct {
	LTS                         CreateFeatureV2ResponseName
	GATEWAY_RESPONSES           CreateFeatureV2ResponseName
	RATELIMIT                   CreateFeatureV2ResponseName
	REQUEST_BODY_SIZE           CreateFeatureV2ResponseName
	BACKEND_TIMEOUT             CreateFeatureV2ResponseName
	APP_TOKEN                   CreateFeatureV2ResponseName
	APP_API_KEY                 CreateFeatureV2ResponseName
	APP_BASIC                   CreateFeatureV2ResponseName
	APP_SECRET                  CreateFeatureV2ResponseName
	APP_JWT                     CreateFeatureV2ResponseName
	PUBLIC_KEY                  CreateFeatureV2ResponseName
	BACKEND_TOKEN_ALLOW         CreateFeatureV2ResponseName
	SIGN_BASIC                  CreateFeatureV2ResponseName
	MULTI_AUTH                  CreateFeatureV2ResponseName
	BACKEND_CLIENT_CERTIFICATE  CreateFeatureV2ResponseName
	SSL_CIPHERS                 CreateFeatureV2ResponseName
	APP_CONFIG                  CreateFeatureV2ResponseName
	GREEN_TUNNEL                CreateFeatureV2ResponseName
	APP_ROUTE                   CreateFeatureV2ResponseName
	DEFAULT_GROUP_HIDE          CreateFeatureV2ResponseName
	CASCADE                     CreateFeatureV2ResponseName
	SANDBOX                     CreateFeatureV2ResponseName
	LIVEDATA_CONFIG             CreateFeatureV2ResponseName
	APICLIENT_FIRST_USE_X_HW_ID CreateFeatureV2ResponseName
	CORS                        CreateFeatureV2ResponseName
	API_TASK                    CreateFeatureV2ResponseName
	APP_QUOTA                   CreateFeatureV2ResponseName
	APP_ACL                     CreateFeatureV2ResponseName
	VPC_BACKUP                  CreateFeatureV2ResponseName
	THROTTLE_STRATEGY           CreateFeatureV2ResponseName
	KAFKA_LOG_PLUGIN_OPTIONS    CreateFeatureV2ResponseName
	GZIP                        CreateFeatureV2ResponseName
	JS_INVOCABLE_CACHE_SIZE     CreateFeatureV2ResponseName
}

func GetCreateFeatureV2ResponseNameEnum() CreateFeatureV2ResponseNameEnum {
	return CreateFeatureV2ResponseNameEnum{
		LTS: CreateFeatureV2ResponseName{
			value: "lts",
		},
		GATEWAY_RESPONSES: CreateFeatureV2ResponseName{
			value: "gateway_responses",
		},
		RATELIMIT: CreateFeatureV2ResponseName{
			value: "ratelimit",
		},
		REQUEST_BODY_SIZE: CreateFeatureV2ResponseName{
			value: "request_body_size",
		},
		BACKEND_TIMEOUT: CreateFeatureV2ResponseName{
			value: "backend_timeout",
		},
		APP_TOKEN: CreateFeatureV2ResponseName{
			value: "app_token",
		},
		APP_API_KEY: CreateFeatureV2ResponseName{
			value: "app_api_key",
		},
		APP_BASIC: CreateFeatureV2ResponseName{
			value: "app_basic",
		},
		APP_SECRET: CreateFeatureV2ResponseName{
			value: "app_secret",
		},
		APP_JWT: CreateFeatureV2ResponseName{
			value: "app_jwt",
		},
		PUBLIC_KEY: CreateFeatureV2ResponseName{
			value: "public_key",
		},
		BACKEND_TOKEN_ALLOW: CreateFeatureV2ResponseName{
			value: "backend_token_allow",
		},
		SIGN_BASIC: CreateFeatureV2ResponseName{
			value: "sign_basic",
		},
		MULTI_AUTH: CreateFeatureV2ResponseName{
			value: "multi_auth",
		},
		BACKEND_CLIENT_CERTIFICATE: CreateFeatureV2ResponseName{
			value: "backend_client_certificate",
		},
		SSL_CIPHERS: CreateFeatureV2ResponseName{
			value: "ssl_ciphers",
		},
		APP_CONFIG: CreateFeatureV2ResponseName{
			value: "app_config",
		},
		GREEN_TUNNEL: CreateFeatureV2ResponseName{
			value: "green_tunnel",
		},
		APP_ROUTE: CreateFeatureV2ResponseName{
			value: "app_route",
		},
		DEFAULT_GROUP_HIDE: CreateFeatureV2ResponseName{
			value: "default_group_hide",
		},
		CASCADE: CreateFeatureV2ResponseName{
			value: "cascade",
		},
		SANDBOX: CreateFeatureV2ResponseName{
			value: "sandbox",
		},
		LIVEDATA_CONFIG: CreateFeatureV2ResponseName{
			value: "livedata_config",
		},
		APICLIENT_FIRST_USE_X_HW_ID: CreateFeatureV2ResponseName{
			value: "apiclient_first_use_x_hw_id",
		},
		CORS: CreateFeatureV2ResponseName{
			value: "cors",
		},
		API_TASK: CreateFeatureV2ResponseName{
			value: "api_task",
		},
		APP_QUOTA: CreateFeatureV2ResponseName{
			value: "app_quota",
		},
		APP_ACL: CreateFeatureV2ResponseName{
			value: "app_acl",
		},
		VPC_BACKUP: CreateFeatureV2ResponseName{
			value: "vpc_backup",
		},
		THROTTLE_STRATEGY: CreateFeatureV2ResponseName{
			value: "throttle_strategy",
		},
		KAFKA_LOG_PLUGIN_OPTIONS: CreateFeatureV2ResponseName{
			value: "kafka_log_plugin_options",
		},
		GZIP: CreateFeatureV2ResponseName{
			value: "gzip",
		},
		JS_INVOCABLE_CACHE_SIZE: CreateFeatureV2ResponseName{
			value: "js_invocable_cache_size",
		},
	}
}

func (c CreateFeatureV2ResponseName) Value() string {
	return c.value
}

func (c CreateFeatureV2ResponseName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFeatureV2ResponseName) UnmarshalJSON(b []byte) error {
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
