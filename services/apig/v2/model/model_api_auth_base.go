package model

import (
	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiAuthBase struct {
	// 授权关系编号

	Id *string `json:"id,omitempty"`
	// API的编号

	ApiId *string `json:"api_id,omitempty"`
	// API的名称

	ApiName *string `json:"api_name,omitempty"`
	// API绑定的分组名称

	GroupName *string `json:"group_name,omitempty"`
	// API类型

	ApiType *int32 `json:"api_type,omitempty"`
	// API的描述信息

	ApiRemark *string `json:"api_remark,omitempty"`
	// api授权绑定的环境ID

	EnvId *string `json:"env_id,omitempty"`
	// 授权者

	AuthRole *string `json:"auth_role,omitempty"`
	// 授权创建的时间

	AuthTime *sdktime.SdkTime `json:"auth_time,omitempty"`
	// APP的名称

	AppName *string `json:"app_name,omitempty"`
	// APP的描述

	AppRemark *string `json:"app_remark,omitempty"`
	// APP的类型：  默认为apig，暂不支持其他类型

	AppType *ApiAuthBaseAppType `json:"app_type,omitempty"`
	// APP的创建者，取值如下： - USER：租户自己创建 - MARKET：API市场分配，暂不支持

	AppCreator *string `json:"app_creator,omitempty"`
	// API的发布编号

	PublishId *string `json:"publish_id,omitempty"`
	// API绑定的分组ID

	GroupId *string `json:"group_id,omitempty"`
	// 授权通道类型 - NORMAL：普通通道 - GREEN：绿色通道  暂不支持，默认NORMAL

	AuthTunnel *ApiAuthBaseAuthTunnel `json:"auth_tunnel,omitempty"`
	// 绿色通道的白名单配置

	AuthWhitelist *[]string `json:"auth_whitelist,omitempty"`
	// 绿色通道的黑名单配置

	AuthBlacklist *[]string `json:"auth_blacklist,omitempty"`
	// 访问参数。

	VisitParam *string `json:"visit_param,omitempty"`
	// ROMA_APP的类型： - subscription：订阅应用 - integration：集成应用  暂不支持

	RomaAppType *string `json:"roma_app_type,omitempty"`
}

func (o ApiAuthBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthBase struct{}"
	}

	return strings.Join([]string{"ApiAuthBase", string(data)}, " ")
}

type ApiAuthBaseAppType struct {
	value string
}

type ApiAuthBaseAppTypeEnum struct {
	APIG ApiAuthBaseAppType
	ROMA ApiAuthBaseAppType
}

func GetApiAuthBaseAppTypeEnum() ApiAuthBaseAppTypeEnum {
	return ApiAuthBaseAppTypeEnum{
		APIG: ApiAuthBaseAppType{
			value: "apig",
		},
		ROMA: ApiAuthBaseAppType{
			value: "roma",
		},
	}
}

func (c ApiAuthBaseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthBaseAppType) UnmarshalJSON(b []byte) error {
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

type ApiAuthBaseAuthTunnel struct {
	value string
}

type ApiAuthBaseAuthTunnelEnum struct {
	NORMAL ApiAuthBaseAuthTunnel
	GREEN  ApiAuthBaseAuthTunnel
}

func GetApiAuthBaseAuthTunnelEnum() ApiAuthBaseAuthTunnelEnum {
	return ApiAuthBaseAuthTunnelEnum{
		NORMAL: ApiAuthBaseAuthTunnel{
			value: "NORMAL",
		},
		GREEN: ApiAuthBaseAuthTunnel{
			value: "GREEN",
		},
	}
}

func (c ApiAuthBaseAuthTunnel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthBaseAuthTunnel) UnmarshalJSON(b []byte) error {
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
