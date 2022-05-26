package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type Config struct {

	// 配额编号
	ConfigId *string `json:"config_id,omitempty"`

	// 配额名称
	ConfigName *ConfigConfigName `json:"config_name,omitempty"`

	// 配额值  当前实例所在租户该配额对应的数量
	ConfigValue *string `json:"config_value,omitempty"`

	// 配额创建时间
	ConfigTime *sdktime.SdkTime `json:"config_time,omitempty"`

	// 配额描述： - API_NUM_LIMIT：租户可以创建的API个数限制 - APP_NUM_LIMIT：租户可以创建的APP个数限制 - APIGROUP_NUM_LIMIT：租户可以创建的API分组个数限制 - ENVIRONMENT_NUM_LIMIT：租户可以创建的环境个数限制 - VARIABLE_NUM_LIMIT：每个API分组上可以创建的环境变量个数限制 - SIGN_NUM_LIMIT：租户可以创建的签名密钥个数限制 - THROTTLE_NUM_LIMIT：租户可以创建的流控策略个数限制 - APIGROUP_DOMAIN_NUM_LIMIT：每个API分组上可以绑定的自定义域名个数限制 - API_VERSION_NUM_LIMIT：每个API可以保留的发布版本个数限制 - VPC_NUM_LIMIT：租户可以创建的VPC通道个数限制 - VPC_INSTANCE_NUM_LIMIT：每个VPC通道上可以绑定的弹性云服务器个数限制 - API_PARAM_NUM_LIMIT：每个API可以设置的参数个数限制 - API_USER_CALL_LIMIT：每个租户的API单位时间内的请求默认限制 - ACL_NUM_LIMIT：每个租户可以创建的ACL策略个数限制 - APP_THROTTLE_LIMIT：特殊应用流控策略个数限制 - USER_THROTTLE_LIMIT：特殊用户流控策略个数限制 - API_NUM_LIMIT_PER_GROUP：租户每个API分组可以创建的API数量限制 - API_POLICY_NUM_LIMIT：每个API可以设置的策略后端个数限制 - API_CONDITION_NUM_LIMIT：每个API策略后端可以设置的条件个数限制 - SL_DOMAIN_CALL_LIMIT：每个二级域名单位时间内的请求默认限制 - ELB_SWITCH：是否启用ELB通道 - AUTHORIZER_NUM_LIMIT：租户可创建的自定义认证个数限制 - AUTHORIZER_IDENTITY_NUM_LIMIT：每个自定义认证可以设置的身份来源个数限制 - APP_CODE_NUM_LIMIT：每个APP可以创建的APP code数量限制 - REGION_MANAGER_WHITELIST_SERVICES：不校验region manager服务白名单列表，暂不支持 - API_SWAGGER_NUM_LIMIT：单个API分组可以绑定的swagger文档数量限制 - API_TAG_NUM_LIMIT：每个API可以设置的标签个数限制 - LTS_SWITCH：是启用LTS上报 - APP_KEY_SECRET_SWITCH：是否打开APP支持自定义KEY和SECRET的开关，1：开启；2：关闭 - RESPONSE_NUM_LIMIT：分组自定义响应个数限制 - CONFIG_NUM_LIMIT_PER_APP：每个APP可以设置的配置项个数限制 - BACKEND_TOKEN_ALLOW_SWITCH：是否支持普通租户透传后端token，1：开启；2：关闭 - APP_TOKEN_SWITCH：是否启用APPTOKEN - API_DESIGNER_SWITCH：是否启用api设计器，1：开启；2：关闭 - APP_API_KEY_SWITCH：是否启用APP_API_KEY认证方式 - APP_BASIC_SWITCH：是否启用APP_BASIC认证方式 - APP_JWT_SWITCH：是否启用APP_JWT认证方式 - APP_ROUTE_SWITCH：是否启用APP路由 - PUBLIC_KEY_SWITCH：是否启用PUBLIK_KEY后端认证方式 - APP_SECRET_SWITCH：是否启用APP_SECRET认证方式 - CASCADE_SWITCH：是否启用级联网关 - IS_INIT_API_PATH_HASH：是否执行过API PATH HASH刷新
	Remark *string `json:"remark,omitempty"`

	// 该实例对应配额已使用数量
	Used *int64 `json:"used,omitempty"`
}

func (o Config) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Config struct{}"
	}

	return strings.Join([]string{"Config", string(data)}, " ")
}

type ConfigConfigName struct {
	value string
}

type ConfigConfigNameEnum struct {
	API_NUM_LIMIT                     ConfigConfigName
	APP_NUM_LIMIT                     ConfigConfigName
	APIGROUP_NUM_LIMIT                ConfigConfigName
	ENVIRONMENT_NUM_LIMIT             ConfigConfigName
	VARIABLE_NUM_LIMIT                ConfigConfigName
	SIGN_NUM_LIMIT                    ConfigConfigName
	THROTTLE_NUM_LIMIT                ConfigConfigName
	APIGROUP_DOMAIN_NUM_LIMIT         ConfigConfigName
	API_VERSION_NUM_LIMIT             ConfigConfigName
	VPC_NUM_LIMIT                     ConfigConfigName
	VPC_INSTANCE_NUM_LIMIT            ConfigConfigName
	API_PARAM_NUM_LIMIT               ConfigConfigName
	API_USER_CALL_LIMIT               ConfigConfigName
	ACL_NUM_LIMIT                     ConfigConfigName
	APP_THROTTLE_LIMIT                ConfigConfigName
	USER_THROTTLE_LIMIT               ConfigConfigName
	API_NUM_LIMIT_PER_GROUP           ConfigConfigName
	API_POLICY_NUM_LIMIT              ConfigConfigName
	API_CONDITION_NUM_LIMIT           ConfigConfigName
	SL_DOMAIN_CALL_LIMIT              ConfigConfigName
	ELB_SWITCH                        ConfigConfigName
	AUTHORIZER_NUM_LIMIT              ConfigConfigName
	AUTHORIZER_IDENTITY_NUM_LIMIT     ConfigConfigName
	APP_CODE_NUM_LIMIT                ConfigConfigName
	REGION_MANAGER_WHITELIST_SERVICES ConfigConfigName
	API_SWAGGER_NUM_LIMIT             ConfigConfigName
	API_TAG_NUM_LIMIT                 ConfigConfigName
	LTS_SWITCH                        ConfigConfigName
	APP_KEY_SECRET_SWITCH             ConfigConfigName
	RESPONSE_NUM_LIMIT                ConfigConfigName
	CONFIG_NUM_LIMIT_PER_APP          ConfigConfigName
	BACKEND_TOKEN_ALLOW_SWITCH        ConfigConfigName
	APP_TOKEN_SWITCH                  ConfigConfigName
	API_DESIGNER_SWITCH               ConfigConfigName
	APP_API_KEY_SWITCH                ConfigConfigName
	APP_BASIC_SWITCH                  ConfigConfigName
	APP_JWT_SWITCH                    ConfigConfigName
	APP_ROUTE_SWITCH                  ConfigConfigName
	PUBLIC_KEY_SWITCH                 ConfigConfigName
	APP_SECRET_SWITCH                 ConfigConfigName
	CASCADE_SWITCH                    ConfigConfigName
	IS_INIT_API_PATH_HASH             ConfigConfigName
}

func GetConfigConfigNameEnum() ConfigConfigNameEnum {
	return ConfigConfigNameEnum{
		API_NUM_LIMIT: ConfigConfigName{
			value: "API_NUM_LIMIT",
		},
		APP_NUM_LIMIT: ConfigConfigName{
			value: "APP_NUM_LIMIT",
		},
		APIGROUP_NUM_LIMIT: ConfigConfigName{
			value: "APIGROUP_NUM_LIMIT",
		},
		ENVIRONMENT_NUM_LIMIT: ConfigConfigName{
			value: "ENVIRONMENT_NUM_LIMIT",
		},
		VARIABLE_NUM_LIMIT: ConfigConfigName{
			value: "VARIABLE_NUM_LIMIT",
		},
		SIGN_NUM_LIMIT: ConfigConfigName{
			value: "SIGN_NUM_LIMIT",
		},
		THROTTLE_NUM_LIMIT: ConfigConfigName{
			value: "THROTTLE_NUM_LIMIT",
		},
		APIGROUP_DOMAIN_NUM_LIMIT: ConfigConfigName{
			value: "APIGROUP_DOMAIN_NUM_LIMIT",
		},
		API_VERSION_NUM_LIMIT: ConfigConfigName{
			value: "API_VERSION_NUM_LIMIT",
		},
		VPC_NUM_LIMIT: ConfigConfigName{
			value: "VPC_NUM_LIMIT",
		},
		VPC_INSTANCE_NUM_LIMIT: ConfigConfigName{
			value: "VPC_INSTANCE_NUM_LIMIT",
		},
		API_PARAM_NUM_LIMIT: ConfigConfigName{
			value: "API_PARAM_NUM_LIMIT",
		},
		API_USER_CALL_LIMIT: ConfigConfigName{
			value: "API_USER_CALL_LIMIT",
		},
		ACL_NUM_LIMIT: ConfigConfigName{
			value: "ACL_NUM_LIMIT",
		},
		APP_THROTTLE_LIMIT: ConfigConfigName{
			value: "APP_THROTTLE_LIMIT",
		},
		USER_THROTTLE_LIMIT: ConfigConfigName{
			value: "USER_THROTTLE_LIMIT",
		},
		API_NUM_LIMIT_PER_GROUP: ConfigConfigName{
			value: "API_NUM_LIMIT_PER_GROUP",
		},
		API_POLICY_NUM_LIMIT: ConfigConfigName{
			value: "API_POLICY_NUM_LIMIT",
		},
		API_CONDITION_NUM_LIMIT: ConfigConfigName{
			value: "API_CONDITION_NUM_LIMIT",
		},
		SL_DOMAIN_CALL_LIMIT: ConfigConfigName{
			value: "SL_DOMAIN_CALL_LIMIT",
		},
		ELB_SWITCH: ConfigConfigName{
			value: "ELB_SWITCH",
		},
		AUTHORIZER_NUM_LIMIT: ConfigConfigName{
			value: "AUTHORIZER_NUM_LIMIT",
		},
		AUTHORIZER_IDENTITY_NUM_LIMIT: ConfigConfigName{
			value: "AUTHORIZER_IDENTITY_NUM_LIMIT",
		},
		APP_CODE_NUM_LIMIT: ConfigConfigName{
			value: "APP_CODE_NUM_LIMIT",
		},
		REGION_MANAGER_WHITELIST_SERVICES: ConfigConfigName{
			value: "REGION_MANAGER_WHITELIST_SERVICES",
		},
		API_SWAGGER_NUM_LIMIT: ConfigConfigName{
			value: "API_SWAGGER_NUM_LIMIT",
		},
		API_TAG_NUM_LIMIT: ConfigConfigName{
			value: "API_TAG_NUM_LIMIT",
		},
		LTS_SWITCH: ConfigConfigName{
			value: "LTS_SWITCH",
		},
		APP_KEY_SECRET_SWITCH: ConfigConfigName{
			value: "APP_KEY_SECRET_SWITCH",
		},
		RESPONSE_NUM_LIMIT: ConfigConfigName{
			value: "RESPONSE_NUM_LIMIT",
		},
		CONFIG_NUM_LIMIT_PER_APP: ConfigConfigName{
			value: "CONFIG_NUM_LIMIT_PER_APP",
		},
		BACKEND_TOKEN_ALLOW_SWITCH: ConfigConfigName{
			value: "BACKEND_TOKEN_ALLOW_SWITCH",
		},
		APP_TOKEN_SWITCH: ConfigConfigName{
			value: "APP_TOKEN_SWITCH",
		},
		API_DESIGNER_SWITCH: ConfigConfigName{
			value: "API_DESIGNER_SWITCH",
		},
		APP_API_KEY_SWITCH: ConfigConfigName{
			value: "APP_API_KEY_SWITCH",
		},
		APP_BASIC_SWITCH: ConfigConfigName{
			value: "APP_BASIC_SWITCH",
		},
		APP_JWT_SWITCH: ConfigConfigName{
			value: "APP_JWT_SWITCH",
		},
		APP_ROUTE_SWITCH: ConfigConfigName{
			value: "APP_ROUTE_SWITCH",
		},
		PUBLIC_KEY_SWITCH: ConfigConfigName{
			value: "PUBLIC_KEY_SWITCH",
		},
		APP_SECRET_SWITCH: ConfigConfigName{
			value: "APP_SECRET_SWITCH",
		},
		CASCADE_SWITCH: ConfigConfigName{
			value: "CASCADE_SWITCH",
		},
		IS_INIT_API_PATH_HASH: ConfigConfigName{
			value: "IS_INIT_API_PATH_HASH",
		},
	}
}

func (c ConfigConfigName) Value() string {
	return c.value
}

func (c ConfigConfigName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigConfigName) UnmarshalJSON(b []byte) error {
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
