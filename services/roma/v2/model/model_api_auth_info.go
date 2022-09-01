package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiAuthInfo struct {

	// 授权关系编号
	Id *string `json:"id,omitempty" xml:"id"`

	// API的编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// API的名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`

	// API绑定的分组名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// API类型
	ApiType *int32 `json:"api_type,omitempty" xml:"api_type"`

	// API的描述信息
	ApiRemark *string `json:"api_remark,omitempty" xml:"api_remark"`

	// api授权绑定的环境ID
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// 授权者
	AuthRole *string `json:"auth_role,omitempty" xml:"auth_role"`

	// 授权创建的时间
	AuthTime *sdktime.SdkTime `json:"auth_time,omitempty" xml:"auth_time"`

	// APP的名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// APP的描述
	AppRemark *string `json:"app_remark,omitempty" xml:"app_remark"`

	// APP的类型： - apig：存量apic客户端，新建实例不支持此类型 - roma：roma集成客户端
	AppType *ApiAuthInfoAppType `json:"app_type,omitempty" xml:"app_type"`

	// APP的创建者，取值如下： - USER：租户自己创建 - MARKET：API市场分配，暂不支持
	AppCreator *string `json:"app_creator,omitempty" xml:"app_creator"`

	// API的发布编号
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`

	// API绑定的分组ID
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 授权通道类型 - NORMAL：普通通道 - GREEN：绿色通道
	AuthTunnel *ApiAuthInfoAuthTunnel `json:"auth_tunnel,omitempty" xml:"auth_tunnel"`

	// 绿色通道的白名单配置
	AuthWhitelist *[]string `json:"auth_whitelist,omitempty" xml:"auth_whitelist"`

	// 绿色通道的黑名单配置
	AuthBlacklist *[]string `json:"auth_blacklist,omitempty" xml:"auth_blacklist"`

	// 访问参数。
	VisitParam *string `json:"visit_param,omitempty" xml:"visit_param"`

	// ROMA_APP的类型： - subscription：订阅应用 - integration：集成应用
	RomaAppType *string `json:"roma_app_type,omitempty" xml:"roma_app_type"`

	// api授权绑定的环境名称
	EnvName *string `json:"env_name,omitempty" xml:"env_name"`

	// APP的编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`
}

func (o ApiAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthInfo struct{}"
	}

	return strings.Join([]string{"ApiAuthInfo", string(data)}, " ")
}

type ApiAuthInfoAppType struct {
	value string
}

type ApiAuthInfoAppTypeEnum struct {
	APIG ApiAuthInfoAppType
	ROMA ApiAuthInfoAppType
}

func GetApiAuthInfoAppTypeEnum() ApiAuthInfoAppTypeEnum {
	return ApiAuthInfoAppTypeEnum{
		APIG: ApiAuthInfoAppType{
			value: "apig",
		},
		ROMA: ApiAuthInfoAppType{
			value: "roma",
		},
	}
}

func (c ApiAuthInfoAppType) Value() string {
	return c.value
}

func (c ApiAuthInfoAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthInfoAppType) UnmarshalJSON(b []byte) error {
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

type ApiAuthInfoAuthTunnel struct {
	value string
}

type ApiAuthInfoAuthTunnelEnum struct {
	NORMAL ApiAuthInfoAuthTunnel
	GREEN  ApiAuthInfoAuthTunnel
}

func GetApiAuthInfoAuthTunnelEnum() ApiAuthInfoAuthTunnelEnum {
	return ApiAuthInfoAuthTunnelEnum{
		NORMAL: ApiAuthInfoAuthTunnel{
			value: "NORMAL",
		},
		GREEN: ApiAuthInfoAuthTunnel{
			value: "GREEN",
		},
	}
}

func (c ApiAuthInfoAuthTunnel) Value() string {
	return c.value
}

func (c ApiAuthInfoAuthTunnel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthInfoAuthTunnel) UnmarshalJSON(b []byte) error {
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
