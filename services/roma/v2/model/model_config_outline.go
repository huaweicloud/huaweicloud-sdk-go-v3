package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ConfigOutline struct {

	// 配额编号
	ConfigId *string `json:"config_id,omitempty"`

	// 配额名称
	ConfigName *ConfigOutlineConfigName `json:"config_name,omitempty"`

	// 配额值  当前实例所在租户该配额对应的数量
	ConfigValue *string `json:"config_value,omitempty"`

	// 配额创建时间
	ConfigTime *sdktime.SdkTime `json:"config_time,omitempty"`

	// 配额描述：   - API_NUM_LIMIT：租户可以创建的API个数限制   - APP_NUM_LIMIT：租户可以创建的APP个数限制   - APIGROUP_NUM_LIMIT：租户可以创建的API分组个数限制   - ENVIRONMENT_NUM_LIMIT：租户可以创建的环境个数限制   - VARIABLE_NUM_LIMIT：每个API分组上可以创建的环境变量个数限制   - SIGN_NUM_LIMIT：租户可以创建的签名密钥个数限制   - THROTTLE_NUM_LIMIT：租户可以创建的流控策略个数限制   - APIGROUP_DOMAIN_NUM_LIMIT：每个API分组上可以绑定的自定义域名个数限制   - API_VERSION_NUM_LIMIT：每个API可以保留的发布版本个数限制   - VPC_NUM_LIMIT：租户可以创建的VPC通道个数限制   - VPC_INSTANCE_NUM_LIMIT：每个VPC通道上可以绑定的弹性云服务器个数限制   - API_PARAM_NUM_LIMIT：每个API可以设置的参数个数限制   - API_USER_CALL_LIMIT：每个租户的API单位时间内的请求默认限制   - ACL_NUM_LIMIT：每个租户可以创建的ACL策略个数限制   - APP_THROTTLE_LIMIT：特殊应用流控策略个数限制   - USER_THROTTLE_LIMIT：特殊用户流控策略个数限制   - API_NUM_LIMIT_PER_GROUP：租户每个API分组可以创建的API数量限制   - API_POLICY_NUM_LIMIT：每个API可以设置的策略后端个数限制   - API_CONDITION_NUM_LIMIT：每个API策略后端可以设置的条件个数限制   - SL_DOMAIN_CALL_LIMIT：每个二级域名单位时间内的请求默认限制   - ELB_SWITCH：是否启用ELB通道   - AUTHORIZER_NUM_LIMIT：租户可创建的自定义认证个数限制   - AUTHORIZER_IDENTITY_NUM_LIMIT：每个自定义认证可以设置的身份来源个数限制   - APP_CODE_NUM_LIMIT：每个APP可以创建的APP code数量限制   - REGION_MANAGER_WHITELIST_SERVICES：不校验region manager服务白名单列表，暂不支持   - API_SWAGGER_NUM_LIMIT：单个API分组可以绑定的swagger文档数量限制   - API_TAG_NUM_LIMIT：每个API可以设置的标签个数限制   - LTS_SWITCH：是启用LTS上报   - APP_KEY_SECRET_SWITCH：是否打开APP支持自定义KEY和SECRET的开关，1：开启；2：关闭   - RESPONSE_NUM_LIMIT：分组自定义响应个数限制   - CONFIG_NUM_LIMIT_PER_APP：每个APP可以设置的配置项个数限制   - BACKEND_TOKEN_ALLOW_SWITCH：是否支持普通租户透传后端token，1：开启；2：关闭   - APP_TOKEN_SWITCH：是否启用APPTOKEN   - API_DESIGNER_SWITCH：是否启用api设计器，1：开启；2：关闭   - APP_API_KEY_SWITCH：是否启用APP_API_KEY认证方式   - APP_BASIC_SWITCH：是否启用APP_BASIC认证方式   - APP_JWT_SWITCH：是否启用APP_JWT认证方式   - APP_ROUTE_SWITCH：是否启用APP路由   - PUBLIC_KEY_SWITCH：是否启用PUBLIK_KEY后端认证方式   - APP_SECRET_SWITCH：是否启用APP_SECRET认证方式   - CASCADE_SWITCH：是否启用级联网关   - IS_INIT_API_PATH_HASH：是否执行过API PATH HASH刷新   - APP_QUOTA_NUM_LIMIT：租户可以创建的客户端配额个数   - IS_INIT_API_VERSION：是否执行过API VERSION CANONICAL PATH刷新   - PLUGIN_NUM_LIMIT：租户可以创建的插件个数   - APICLIENT_FIRST_USE_X_HW_ID_SWITCH：ApiClient是否优先使用x-hw-id校验权限   - API_TASK_NUM_LIMIT：[租户可以创建的API定时任务个数限制](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - THROTTLE_LOCAL_SWITCH：[是否启用本地流控模式](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - API_TASK_SWITCH：[租户是否支持定时任务](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - SET_HEADERS_NUM_LIMIT_PER_PLUGIN：[租户可以通过插件创建的HTTP头个数限制](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - LUA_SCRIPT_SWITCH：[租户是否允许使用lua_script插件](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - LUA_SCRIPT_NUM_LIMIT：[每个实例可以创建的lua_script类型插件个数限制](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - SM_DICT_NUM_LIMIT：[每个实例可以创建的数据字典个数限制](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)   - BM_VPC_INSTANCE_GROUP_NUM_LIMIT：[每个实例可以创建的VPC通道后端服务器组个数限制](tag:hws,hws_hk)[暂未使用](tag:fcs,hcs,hcs_sm,g42,Site)
	Remark *string `json:"remark,omitempty"`
}

func (o ConfigOutline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigOutline struct{}"
	}

	return strings.Join([]string{"ConfigOutline", string(data)}, " ")
}

type ConfigOutlineConfigName struct {
	value string
}

type ConfigOutlineConfigNameEnum struct {
	API_NUM_LIMIT                      ConfigOutlineConfigName
	APP_NUM_LIMIT                      ConfigOutlineConfigName
	APIGROUP_NUM_LIMIT                 ConfigOutlineConfigName
	ENVIRONMENT_NUM_LIMIT              ConfigOutlineConfigName
	VARIABLE_NUM_LIMIT                 ConfigOutlineConfigName
	SIGN_NUM_LIMIT                     ConfigOutlineConfigName
	THROTTLE_NUM_LIMIT                 ConfigOutlineConfigName
	APIGROUP_DOMAIN_NUM_LIMIT          ConfigOutlineConfigName
	API_VERSION_NUM_LIMIT              ConfigOutlineConfigName
	VPC_NUM_LIMIT                      ConfigOutlineConfigName
	VPC_INSTANCE_NUM_LIMIT             ConfigOutlineConfigName
	API_PARAM_NUM_LIMIT                ConfigOutlineConfigName
	API_USER_CALL_LIMIT                ConfigOutlineConfigName
	ACL_NUM_LIMIT                      ConfigOutlineConfigName
	APP_THROTTLE_LIMIT                 ConfigOutlineConfigName
	USER_THROTTLE_LIMIT                ConfigOutlineConfigName
	API_NUM_LIMIT_PER_GROUP            ConfigOutlineConfigName
	API_POLICY_NUM_LIMIT               ConfigOutlineConfigName
	API_CONDITION_NUM_LIMIT            ConfigOutlineConfigName
	SL_DOMAIN_CALL_LIMIT               ConfigOutlineConfigName
	ELB_SWITCH                         ConfigOutlineConfigName
	AUTHORIZER_NUM_LIMIT               ConfigOutlineConfigName
	AUTHORIZER_IDENTITY_NUM_LIMIT      ConfigOutlineConfigName
	APP_CODE_NUM_LIMIT                 ConfigOutlineConfigName
	REGION_MANAGER_WHITELIST_SERVICES  ConfigOutlineConfigName
	API_SWAGGER_NUM_LIMIT              ConfigOutlineConfigName
	API_TAG_NUM_LIMIT                  ConfigOutlineConfigName
	LTS_SWITCH                         ConfigOutlineConfigName
	APP_KEY_SECRET_SWITCH              ConfigOutlineConfigName
	RESPONSE_NUM_LIMIT                 ConfigOutlineConfigName
	CONFIG_NUM_LIMIT_PER_APP           ConfigOutlineConfigName
	BACKEND_TOKEN_ALLOW_SWITCH         ConfigOutlineConfigName
	APP_TOKEN_SWITCH                   ConfigOutlineConfigName
	API_DESIGNER_SWITCH                ConfigOutlineConfigName
	APP_API_KEY_SWITCH                 ConfigOutlineConfigName
	APP_BASIC_SWITCH                   ConfigOutlineConfigName
	APP_JWT_SWITCH                     ConfigOutlineConfigName
	APP_ROUTE_SWITCH                   ConfigOutlineConfigName
	PUBLIC_KEY_SWITCH                  ConfigOutlineConfigName
	APP_SECRET_SWITCH                  ConfigOutlineConfigName
	CASCADE_SWITCH                     ConfigOutlineConfigName
	IS_INIT_API_PATH_HASH              ConfigOutlineConfigName
	APP_QUOTA_NUM_LIMIT                ConfigOutlineConfigName
	IS_INIT_API_VERSION                ConfigOutlineConfigName
	PLUGIN_NUM_LIMIT                   ConfigOutlineConfigName
	APICLIENT_FIRST_USE_X_HW_ID_SWITCH ConfigOutlineConfigName
	API_TASK_NUM_LIMIT                 ConfigOutlineConfigName
	SET_HEADERS_NUM_LIMIT_PER_PLUGIN   ConfigOutlineConfigName
	API_TASK_SWITCH                    ConfigOutlineConfigName
	THROTTLE_LOCAL_SWITCH              ConfigOutlineConfigName
	LUA_SCRIPT_SWITCH                  ConfigOutlineConfigName
	SM_DICT_NUM_LIMIT                  ConfigOutlineConfigName
	BM_VPC_INSTANCE_GROUP_NUM_LIMIT    ConfigOutlineConfigName
}

func GetConfigOutlineConfigNameEnum() ConfigOutlineConfigNameEnum {
	return ConfigOutlineConfigNameEnum{
		API_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_NUM_LIMIT",
		},
		APP_NUM_LIMIT: ConfigOutlineConfigName{
			value: "APP_NUM_LIMIT",
		},
		APIGROUP_NUM_LIMIT: ConfigOutlineConfigName{
			value: "APIGROUP_NUM_LIMIT",
		},
		ENVIRONMENT_NUM_LIMIT: ConfigOutlineConfigName{
			value: "ENVIRONMENT_NUM_LIMIT",
		},
		VARIABLE_NUM_LIMIT: ConfigOutlineConfigName{
			value: "VARIABLE_NUM_LIMIT",
		},
		SIGN_NUM_LIMIT: ConfigOutlineConfigName{
			value: "SIGN_NUM_LIMIT",
		},
		THROTTLE_NUM_LIMIT: ConfigOutlineConfigName{
			value: "THROTTLE_NUM_LIMIT",
		},
		APIGROUP_DOMAIN_NUM_LIMIT: ConfigOutlineConfigName{
			value: "APIGROUP_DOMAIN_NUM_LIMIT",
		},
		API_VERSION_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_VERSION_NUM_LIMIT",
		},
		VPC_NUM_LIMIT: ConfigOutlineConfigName{
			value: "VPC_NUM_LIMIT",
		},
		VPC_INSTANCE_NUM_LIMIT: ConfigOutlineConfigName{
			value: "VPC_INSTANCE_NUM_LIMIT",
		},
		API_PARAM_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_PARAM_NUM_LIMIT",
		},
		API_USER_CALL_LIMIT: ConfigOutlineConfigName{
			value: "API_USER_CALL_LIMIT",
		},
		ACL_NUM_LIMIT: ConfigOutlineConfigName{
			value: "ACL_NUM_LIMIT",
		},
		APP_THROTTLE_LIMIT: ConfigOutlineConfigName{
			value: "APP_THROTTLE_LIMIT",
		},
		USER_THROTTLE_LIMIT: ConfigOutlineConfigName{
			value: "USER_THROTTLE_LIMIT",
		},
		API_NUM_LIMIT_PER_GROUP: ConfigOutlineConfigName{
			value: "API_NUM_LIMIT_PER_GROUP",
		},
		API_POLICY_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_POLICY_NUM_LIMIT",
		},
		API_CONDITION_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_CONDITION_NUM_LIMIT",
		},
		SL_DOMAIN_CALL_LIMIT: ConfigOutlineConfigName{
			value: "SL_DOMAIN_CALL_LIMIT",
		},
		ELB_SWITCH: ConfigOutlineConfigName{
			value: "ELB_SWITCH",
		},
		AUTHORIZER_NUM_LIMIT: ConfigOutlineConfigName{
			value: "AUTHORIZER_NUM_LIMIT",
		},
		AUTHORIZER_IDENTITY_NUM_LIMIT: ConfigOutlineConfigName{
			value: "AUTHORIZER_IDENTITY_NUM_LIMIT",
		},
		APP_CODE_NUM_LIMIT: ConfigOutlineConfigName{
			value: "APP_CODE_NUM_LIMIT",
		},
		REGION_MANAGER_WHITELIST_SERVICES: ConfigOutlineConfigName{
			value: "REGION_MANAGER_WHITELIST_SERVICES",
		},
		API_SWAGGER_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_SWAGGER_NUM_LIMIT",
		},
		API_TAG_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_TAG_NUM_LIMIT",
		},
		LTS_SWITCH: ConfigOutlineConfigName{
			value: "LTS_SWITCH",
		},
		APP_KEY_SECRET_SWITCH: ConfigOutlineConfigName{
			value: "APP_KEY_SECRET_SWITCH",
		},
		RESPONSE_NUM_LIMIT: ConfigOutlineConfigName{
			value: "RESPONSE_NUM_LIMIT",
		},
		CONFIG_NUM_LIMIT_PER_APP: ConfigOutlineConfigName{
			value: "CONFIG_NUM_LIMIT_PER_APP",
		},
		BACKEND_TOKEN_ALLOW_SWITCH: ConfigOutlineConfigName{
			value: "BACKEND_TOKEN_ALLOW_SWITCH",
		},
		APP_TOKEN_SWITCH: ConfigOutlineConfigName{
			value: "APP_TOKEN_SWITCH",
		},
		API_DESIGNER_SWITCH: ConfigOutlineConfigName{
			value: "API_DESIGNER_SWITCH",
		},
		APP_API_KEY_SWITCH: ConfigOutlineConfigName{
			value: "APP_API_KEY_SWITCH",
		},
		APP_BASIC_SWITCH: ConfigOutlineConfigName{
			value: "APP_BASIC_SWITCH",
		},
		APP_JWT_SWITCH: ConfigOutlineConfigName{
			value: "APP_JWT_SWITCH",
		},
		APP_ROUTE_SWITCH: ConfigOutlineConfigName{
			value: "APP_ROUTE_SWITCH",
		},
		PUBLIC_KEY_SWITCH: ConfigOutlineConfigName{
			value: "PUBLIC_KEY_SWITCH",
		},
		APP_SECRET_SWITCH: ConfigOutlineConfigName{
			value: "APP_SECRET_SWITCH",
		},
		CASCADE_SWITCH: ConfigOutlineConfigName{
			value: "CASCADE_SWITCH",
		},
		IS_INIT_API_PATH_HASH: ConfigOutlineConfigName{
			value: "IS_INIT_API_PATH_HASH",
		},
		APP_QUOTA_NUM_LIMIT: ConfigOutlineConfigName{
			value: "APP_QUOTA_NUM_LIMIT",
		},
		IS_INIT_API_VERSION: ConfigOutlineConfigName{
			value: "IS_INIT_API_VERSION",
		},
		PLUGIN_NUM_LIMIT: ConfigOutlineConfigName{
			value: "PLUGIN_NUM_LIMIT",
		},
		APICLIENT_FIRST_USE_X_HW_ID_SWITCH: ConfigOutlineConfigName{
			value: "APICLIENT_FIRST_USE_X_HW_ID_SWITCH",
		},
		API_TASK_NUM_LIMIT: ConfigOutlineConfigName{
			value: "API_TASK_NUM_LIMIT",
		},
		SET_HEADERS_NUM_LIMIT_PER_PLUGIN: ConfigOutlineConfigName{
			value: "SET_HEADERS_NUM_LIMIT_PER_PLUGIN",
		},
		API_TASK_SWITCH: ConfigOutlineConfigName{
			value: "API_TASK_SWITCH",
		},
		THROTTLE_LOCAL_SWITCH: ConfigOutlineConfigName{
			value: "THROTTLE_LOCAL_SWITCH",
		},
		LUA_SCRIPT_SWITCH: ConfigOutlineConfigName{
			value: "LUA_SCRIPT_SWITCH",
		},
		SM_DICT_NUM_LIMIT: ConfigOutlineConfigName{
			value: "SM_DICT_NUM_LIMIT",
		},
		BM_VPC_INSTANCE_GROUP_NUM_LIMIT: ConfigOutlineConfigName{
			value: "BM_VPC_INSTANCE_GROUP_NUM_LIMIT",
		},
	}
}

func (c ConfigOutlineConfigName) Value() string {
	return c.value
}

func (c ConfigOutlineConfigName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigOutlineConfigName) UnmarshalJSON(b []byte) error {
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
