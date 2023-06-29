package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LdApiDeploy struct {

	// 是否自动发布API - true：部署完成后自动创建并发布前端API。此时auth_type，group_id，env_id，protocol必填。 - false：部署完成后不创建前端API
	DeployFrontApi *bool `json:"deploy_front_api,omitempty"`

	// 认证方式[，site暂不支持IAM认证。](tag:Site) - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *string `json:"auth_type,omitempty"`

	// 自定义认证编号。  认证方式auth_type = AUTHORIZER时必填
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 部署的前端API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 部署的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// 请求协议
	Protocol *LdApiDeployProtocol `json:"protocol,omitempty"`

	// 超时时间
	BackendTimeout *int32 `json:"backend_timeout,omitempty"`

	// 请求路径
	Path *string `json:"path,omitempty"`

	// 请求方式
	Method *string `json:"method,omitempty"`

	// 是否支持跨域 - true：支持 - false：不支持
	Cors *bool `json:"cors,omitempty"`

	// 部署到前端的api归属的应用编号，与后端归属的应用编号保持一致
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// ROMA Connect APIC请求后端服务的重试次数，默认为-1，范围[-1,10]
	RetryCount *string `json:"retry_count,omitempty"`
}

func (o LdApiDeploy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiDeploy struct{}"
	}

	return strings.Join([]string{"LdApiDeploy", string(data)}, " ")
}

type LdApiDeployProtocol struct {
	value string
}

type LdApiDeployProtocolEnum struct {
	HTTPS     LdApiDeployProtocol
	HTTP      LdApiDeployProtocol
	HTTPSHTTP LdApiDeployProtocol
}

func GetLdApiDeployProtocolEnum() LdApiDeployProtocolEnum {
	return LdApiDeployProtocolEnum{
		HTTPS: LdApiDeployProtocol{
			value: "HTTPS",
		},
		HTTP: LdApiDeployProtocol{
			value: "HTTP",
		},
		HTTPSHTTP: LdApiDeployProtocol{
			value: "HTTPS&HTTP",
		},
	}
}

func (c LdApiDeployProtocol) Value() string {
	return c.value
}

func (c LdApiDeployProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiDeployProtocol) UnmarshalJSON(b []byte) error {
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
