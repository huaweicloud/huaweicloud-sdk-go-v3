package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiAuthRelations struct {

	// API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	AuthResult *AuthResult `json:"auth_result,omitempty" xml:"auth_result"`

	// 授权时间
	AuthTime *sdktime.SdkTime `json:"auth_time,omitempty" xml:"auth_time"`

	// 授权关系编号
	Id *string `json:"id,omitempty" xml:"id"`

	// APP编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 授权者 - PROVIDER：API提供者授权 - CONSUMER：API消费者授权
	AuthRole *ApiAuthRelationsAuthRole `json:"auth_role,omitempty" xml:"auth_role"`

	// 授权通道类型 - NORMAL：普通通道 - GREEN：绿色通道  暂不支持，默认NORMAL
	AuthTunnel *ApiAuthRelationsAuthTunnel `json:"auth_tunnel,omitempty" xml:"auth_tunnel"`

	// 绿色通道的白名单配置
	AuthWhitelist *[]string `json:"auth_whitelist,omitempty" xml:"auth_whitelist"`

	// 绿色通道的黑名单配置
	AuthBlacklist *[]string `json:"auth_blacklist,omitempty" xml:"auth_blacklist"`

	// 访问参数。
	VisitParams *string `json:"visit_params,omitempty" xml:"visit_params"`
}

func (o ApiAuthRelations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthRelations struct{}"
	}

	return strings.Join([]string{"ApiAuthRelations", string(data)}, " ")
}

type ApiAuthRelationsAuthRole struct {
	value string
}

type ApiAuthRelationsAuthRoleEnum struct {
	PROVIDER ApiAuthRelationsAuthRole
	CONSUMER ApiAuthRelationsAuthRole
}

func GetApiAuthRelationsAuthRoleEnum() ApiAuthRelationsAuthRoleEnum {
	return ApiAuthRelationsAuthRoleEnum{
		PROVIDER: ApiAuthRelationsAuthRole{
			value: "PROVIDER",
		},
		CONSUMER: ApiAuthRelationsAuthRole{
			value: "CONSUMER",
		},
	}
}

func (c ApiAuthRelationsAuthRole) Value() string {
	return c.value
}

func (c ApiAuthRelationsAuthRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthRelationsAuthRole) UnmarshalJSON(b []byte) error {
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

type ApiAuthRelationsAuthTunnel struct {
	value string
}

type ApiAuthRelationsAuthTunnelEnum struct {
	NORMAL ApiAuthRelationsAuthTunnel
	GREEN  ApiAuthRelationsAuthTunnel
}

func GetApiAuthRelationsAuthTunnelEnum() ApiAuthRelationsAuthTunnelEnum {
	return ApiAuthRelationsAuthTunnelEnum{
		NORMAL: ApiAuthRelationsAuthTunnel{
			value: "NORMAL",
		},
		GREEN: ApiAuthRelationsAuthTunnel{
			value: "GREEN",
		},
	}
}

func (c ApiAuthRelationsAuthTunnel) Value() string {
	return c.value
}

func (c ApiAuthRelationsAuthTunnel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAuthRelationsAuthTunnel) UnmarshalJSON(b []byte) error {
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
