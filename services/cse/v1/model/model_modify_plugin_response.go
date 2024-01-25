package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyPluginResponse Response Object
type ModifyPluginResponse struct {

	// 拉取OCI镜像的行为
	ImagePullPolicy *ModifyPluginResponseImagePullPolicy `json:"imagePullPolicy,omitempty"`

	// 拉取OCI 镜像的凭据
	ImagePullSecret *string `json:"imagePullSecret,omitempty"`

	// 确定插件将在过滤器链中的何处注入。
	Phase *ModifyPluginResponsePhase `json:"phase,omitempty"`

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

func (o ModifyPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPluginResponse struct{}"
	}

	return strings.Join([]string{"ModifyPluginResponse", string(data)}, " ")
}

type ModifyPluginResponseImagePullPolicy struct {
	value string
}

type ModifyPluginResponseImagePullPolicyEnum struct {
	UNSPECIFIED_POLICY ModifyPluginResponseImagePullPolicy
	IF_NOT_PRESENT     ModifyPluginResponseImagePullPolicy
	ALWAYS             ModifyPluginResponseImagePullPolicy
}

func GetModifyPluginResponseImagePullPolicyEnum() ModifyPluginResponseImagePullPolicyEnum {
	return ModifyPluginResponseImagePullPolicyEnum{
		UNSPECIFIED_POLICY: ModifyPluginResponseImagePullPolicy{
			value: "UNSPECIFIED_POLICY",
		},
		IF_NOT_PRESENT: ModifyPluginResponseImagePullPolicy{
			value: "IfNotPresent",
		},
		ALWAYS: ModifyPluginResponseImagePullPolicy{
			value: "Always",
		},
	}
}

func (c ModifyPluginResponseImagePullPolicy) Value() string {
	return c.value
}

func (c ModifyPluginResponseImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyPluginResponseImagePullPolicy) UnmarshalJSON(b []byte) error {
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

type ModifyPluginResponsePhase struct {
	value string
}

type ModifyPluginResponsePhaseEnum struct {
	UNSPECIFIED_PHASE ModifyPluginResponsePhase
	AUTHN             ModifyPluginResponsePhase
	AUTHZ             ModifyPluginResponsePhase
	STATS             ModifyPluginResponsePhase
}

func GetModifyPluginResponsePhaseEnum() ModifyPluginResponsePhaseEnum {
	return ModifyPluginResponsePhaseEnum{
		UNSPECIFIED_PHASE: ModifyPluginResponsePhase{
			value: "UNSPECIFIED_PHASE",
		},
		AUTHN: ModifyPluginResponsePhase{
			value: "AUTHN",
		},
		AUTHZ: ModifyPluginResponsePhase{
			value: "AUTHZ",
		},
		STATS: ModifyPluginResponsePhase{
			value: "STATS",
		},
	}
}

func (c ModifyPluginResponsePhase) Value() string {
	return c.value
}

func (c ModifyPluginResponsePhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyPluginResponsePhase) UnmarshalJSON(b []byte) error {
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
