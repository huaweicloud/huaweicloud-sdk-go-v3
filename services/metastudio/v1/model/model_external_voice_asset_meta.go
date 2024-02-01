package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExternalVoiceAssetMeta 第三方TTS音色元数据。
type ExternalVoiceAssetMeta struct {

	// 第三方TTS供应商类型。 * XIMALAYA：喜马拉雅TTS * HUAWEI_EI: EI TTS * MOBVOI: 出门问问 TTSS
	Provider ExternalVoiceAssetMetaProvider `json:"provider"`

	MobvoiVoiceMeta *MobvoiVoiceAssetMeta `json:"mobvoi_voice_meta,omitempty"`
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
	MOBVOI    ExternalVoiceAssetMetaProvider
}

func GetExternalVoiceAssetMetaProviderEnum() ExternalVoiceAssetMetaProviderEnum {
	return ExternalVoiceAssetMetaProviderEnum{
		XIMALAYA: ExternalVoiceAssetMetaProvider{
			value: "XIMALAYA",
		},
		HUAWEI_EI: ExternalVoiceAssetMetaProvider{
			value: "HUAWEI_EI",
		},
		MOBVOI: ExternalVoiceAssetMetaProvider{
			value: "MOBVOI",
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
