package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExternalVoiceAssetMeta 第三方TTS音色元数据。此参数仅内部使用，不对外开放。
type ExternalVoiceAssetMeta struct {

	// 第三方TTS供应商类型。 * XIMALAYA：喜马拉雅TTS * HUAWEI_EI: 华为云EI TTS
	Provider ExternalVoiceAssetMetaProvider `json:"provider"`

	XimalayaVoiceMeta *XimalayaVoiceAssetMeta `json:"ximalaya_voice_meta,omitempty"`

	HuaweiEiVoiceMeta *HuaweiEiVoiceAssetMeta `json:"huawei_ei_voice_meta,omitempty"`
}

func (o ExternalVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"ExternalVoiceAssetMeta", string(data)}, " ")
}

type ExternalVoiceAssetMetaProvider struct {
	value string
}

type ExternalVoiceAssetMetaProviderEnum struct {
	XIMALAYA  ExternalVoiceAssetMetaProvider
	HUAWEI_EI ExternalVoiceAssetMetaProvider
}

func GetExternalVoiceAssetMetaProviderEnum() ExternalVoiceAssetMetaProviderEnum {
	return ExternalVoiceAssetMetaProviderEnum{
		XIMALAYA: ExternalVoiceAssetMetaProvider{
			value: "XIMALAYA",
		},
		HUAWEI_EI: ExternalVoiceAssetMetaProvider{
			value: "HUAWEI_EI",
		},
	}
}

func (c ExternalVoiceAssetMetaProvider) Value() string {
	return c.value
}

func (c ExternalVoiceAssetMetaProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExternalVoiceAssetMetaProvider) UnmarshalJSON(b []byte) error {
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
