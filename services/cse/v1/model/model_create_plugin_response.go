package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePluginResponse Response Object
type CreatePluginResponse struct {

	// 拉取OCI镜像的行为
	ImagePullPolicy *CreatePluginResponseImagePullPolicy `json:"imagePullPolicy,omitempty"`

	// 拉取OCI 镜像的凭据
	ImagePullSecret *string `json:"imagePullSecret,omitempty"`

	// 确定插件将在过滤器链中的何处注入。
	Phase *CreatePluginResponsePhase `json:"phase,omitempty"`

	// 传递给插件的配置。
	PluginConfig *interface{} `json:"pluginConfig,omitempty"`

	// 插件名。
	PluginName *string `json:"pluginName,omitempty"`

	// 插件的调用优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 用于校验插件和容器的校验和。
	Sha256 *string `json:"sha256,omitempty"`

	// 插件或容器的下载地址。
	Url *string `json:"url,omitempty"`

	// 校验值。
	VerificationKey *string `json:"verificationKey,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreatePluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginResponse struct{}"
	}

	return strings.Join([]string{"CreatePluginResponse", string(data)}, " ")
}

type CreatePluginResponseImagePullPolicy struct {
	value string
}

type CreatePluginResponseImagePullPolicyEnum struct {
	UNSPECIFIED_POLICY CreatePluginResponseImagePullPolicy
	IF_NOT_PRESENT     CreatePluginResponseImagePullPolicy
	ALWAYS             CreatePluginResponseImagePullPolicy
}

func GetCreatePluginResponseImagePullPolicyEnum() CreatePluginResponseImagePullPolicyEnum {
	return CreatePluginResponseImagePullPolicyEnum{
		UNSPECIFIED_POLICY: CreatePluginResponseImagePullPolicy{
			value: "UNSPECIFIED_POLICY",
		},
		IF_NOT_PRESENT: CreatePluginResponseImagePullPolicy{
			value: "IfNotPresent",
		},
		ALWAYS: CreatePluginResponseImagePullPolicy{
			value: "Always",
		},
	}
}

func (c CreatePluginResponseImagePullPolicy) Value() string {
	return c.value
}

func (c CreatePluginResponseImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePluginResponseImagePullPolicy) UnmarshalJSON(b []byte) error {
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

type CreatePluginResponsePhase struct {
	value string
}

type CreatePluginResponsePhaseEnum struct {
	UNSPECIFIED_PHASE CreatePluginResponsePhase
	AUTHN             CreatePluginResponsePhase
	AUTHZ             CreatePluginResponsePhase
	STATS             CreatePluginResponsePhase
}

func GetCreatePluginResponsePhaseEnum() CreatePluginResponsePhaseEnum {
	return CreatePluginResponsePhaseEnum{
		UNSPECIFIED_PHASE: CreatePluginResponsePhase{
			value: "UNSPECIFIED_PHASE",
		},
		AUTHN: CreatePluginResponsePhase{
			value: "AUTHN",
		},
		AUTHZ: CreatePluginResponsePhase{
			value: "AUTHZ",
		},
		STATS: CreatePluginResponsePhase{
			value: "STATS",
		},
	}
}

func (c CreatePluginResponsePhase) Value() string {
	return c.value
}

func (c CreatePluginResponsePhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePluginResponsePhase) UnmarshalJSON(b []byte) error {
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
