package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSinglePluginResponse Response Object
type ShowSinglePluginResponse struct {

	// 拉取OCI镜像的行为
	ImagePullPolicy *ShowSinglePluginResponseImagePullPolicy `json:"imagePullPolicy,omitempty"`

	// 拉取OCI 镜像的凭据
	ImagePullSecret *string `json:"imagePullSecret,omitempty"`

	// 确定插件将在过滤器链中的何处注入。
	Phase *ShowSinglePluginResponsePhase `json:"phase,omitempty"`

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

func (o ShowSinglePluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinglePluginResponse struct{}"
	}

	return strings.Join([]string{"ShowSinglePluginResponse", string(data)}, " ")
}

type ShowSinglePluginResponseImagePullPolicy struct {
	value string
}

type ShowSinglePluginResponseImagePullPolicyEnum struct {
	UNSPECIFIED_POLICY ShowSinglePluginResponseImagePullPolicy
	IF_NOT_PRESENT     ShowSinglePluginResponseImagePullPolicy
	ALWAYS             ShowSinglePluginResponseImagePullPolicy
}

func GetShowSinglePluginResponseImagePullPolicyEnum() ShowSinglePluginResponseImagePullPolicyEnum {
	return ShowSinglePluginResponseImagePullPolicyEnum{
		UNSPECIFIED_POLICY: ShowSinglePluginResponseImagePullPolicy{
			value: "UNSPECIFIED_POLICY",
		},
		IF_NOT_PRESENT: ShowSinglePluginResponseImagePullPolicy{
			value: "IfNotPresent",
		},
		ALWAYS: ShowSinglePluginResponseImagePullPolicy{
			value: "Always",
		},
	}
}

func (c ShowSinglePluginResponseImagePullPolicy) Value() string {
	return c.value
}

func (c ShowSinglePluginResponseImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSinglePluginResponseImagePullPolicy) UnmarshalJSON(b []byte) error {
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

type ShowSinglePluginResponsePhase struct {
	value string
}

type ShowSinglePluginResponsePhaseEnum struct {
	UNSPECIFIED_PHASE ShowSinglePluginResponsePhase
	AUTHN             ShowSinglePluginResponsePhase
	AUTHZ             ShowSinglePluginResponsePhase
	STATS             ShowSinglePluginResponsePhase
}

func GetShowSinglePluginResponsePhaseEnum() ShowSinglePluginResponsePhaseEnum {
	return ShowSinglePluginResponsePhaseEnum{
		UNSPECIFIED_PHASE: ShowSinglePluginResponsePhase{
			value: "UNSPECIFIED_PHASE",
		},
		AUTHN: ShowSinglePluginResponsePhase{
			value: "AUTHN",
		},
		AUTHZ: ShowSinglePluginResponsePhase{
			value: "AUTHZ",
		},
		STATS: ShowSinglePluginResponsePhase{
			value: "STATS",
		},
	}
}

func (c ShowSinglePluginResponsePhase) Value() string {
	return c.value
}

func (c ShowSinglePluginResponsePhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSinglePluginResponsePhase) UnmarshalJSON(b []byte) error {
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
