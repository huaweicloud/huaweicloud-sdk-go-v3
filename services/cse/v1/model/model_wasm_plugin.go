package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WasmPlugin struct {

	// 拉取OCI镜像的行为
	ImagePullPolicy *WasmPluginImagePullPolicy `json:"imagePullPolicy,omitempty"`

	// 拉取OCI 镜像的凭据
	ImagePullSecret *string `json:"imagePullSecret,omitempty"`

	// 确定插件将在过滤器链中的何处注入。
	Phase *WasmPluginPhase `json:"phase,omitempty"`

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
}

func (o WasmPlugin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WasmPlugin struct{}"
	}

	return strings.Join([]string{"WasmPlugin", string(data)}, " ")
}

type WasmPluginImagePullPolicy struct {
	value string
}

type WasmPluginImagePullPolicyEnum struct {
	UNSPECIFIED_POLICY WasmPluginImagePullPolicy
	IF_NOT_PRESENT     WasmPluginImagePullPolicy
	ALWAYS             WasmPluginImagePullPolicy
}

func GetWasmPluginImagePullPolicyEnum() WasmPluginImagePullPolicyEnum {
	return WasmPluginImagePullPolicyEnum{
		UNSPECIFIED_POLICY: WasmPluginImagePullPolicy{
			value: "UNSPECIFIED_POLICY",
		},
		IF_NOT_PRESENT: WasmPluginImagePullPolicy{
			value: "IfNotPresent",
		},
		ALWAYS: WasmPluginImagePullPolicy{
			value: "Always",
		},
	}
}

func (c WasmPluginImagePullPolicy) Value() string {
	return c.value
}

func (c WasmPluginImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WasmPluginImagePullPolicy) UnmarshalJSON(b []byte) error {
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

type WasmPluginPhase struct {
	value string
}

type WasmPluginPhaseEnum struct {
	UNSPECIFIED_PHASE WasmPluginPhase
	AUTHN             WasmPluginPhase
	AUTHZ             WasmPluginPhase
	STATS             WasmPluginPhase
}

func GetWasmPluginPhaseEnum() WasmPluginPhaseEnum {
	return WasmPluginPhaseEnum{
		UNSPECIFIED_PHASE: WasmPluginPhase{
			value: "UNSPECIFIED_PHASE",
		},
		AUTHN: WasmPluginPhase{
			value: "AUTHN",
		},
		AUTHZ: WasmPluginPhase{
			value: "AUTHZ",
		},
		STATS: WasmPluginPhase{
			value: "STATS",
		},
	}
}

func (c WasmPluginPhase) Value() string {
	return c.value
}

func (c WasmPluginPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WasmPluginPhase) UnmarshalJSON(b []byte) error {
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
