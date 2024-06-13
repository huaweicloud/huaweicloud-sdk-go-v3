package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExternalVoiceAssetMeta 第三方TTS音色元数据。
type ExternalVoiceAssetMeta struct {

	// 第三方TTS供应商类型。 * XIMALAYA：喜马拉雅TTS * HUAWEI_EI: 华为云EI TTS * MOBVOI: 出门问问 TTSS * AUDIOX: 逻辑智能 TTSS * SINOVOICE: 捷通华声 TTSS * DATABACK: 标贝 TTSS * AISPEECH: 思必驰 TTSS
	Provider ExternalVoiceAssetMetaProvider `json:"provider"`
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
	AUDIOX    ExternalVoiceAssetMetaProvider
	SINOVOICE ExternalVoiceAssetMetaProvider
	DATABACK  ExternalVoiceAssetMetaProvider
	AISPEECH  ExternalVoiceAssetMetaProvider
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
		AUDIOX: ExternalVoiceAssetMetaProvider{
			value: "AUDIOX",
		},
		SINOVOICE: ExternalVoiceAssetMetaProvider{
			value: "SINOVOICE",
		},
		DATABACK: ExternalVoiceAssetMetaProvider{
			value: "DATABACK",
		},
		AISPEECH: ExternalVoiceAssetMetaProvider{
			value: "AISPEECH",
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
