package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Encryption 加密信息
type Encryption struct {

	// 密钥缓存时间。如果密钥不变，默认缓存七天。  请注意：目前为保留字段，不支持配置。
	KeyRotationIntervalSeconds *int32 `json:"key_rotation_interval_seconds,omitempty"`

	// 加密方式。  请注意：目前为保留字段，不支持配置。
	EncryptionMethod *EncryptionEncryptionMethod `json:"encryption_method,omitempty"`

	// 取值如下： - content：一个频道对应一个密钥 - profile：一个码率对应一个密钥  默认值：content
	Level *EncryptionLevel `json:"level,omitempty"`

	// 客户生成的DRM内容ID
	ResourceId string `json:"resource_id"`

	// system_id枚举值。  取值如下： * HLS：FairPlay * DASH：Widevine、PlayReady * MSS：PlayReady
	SystemIds []EncryptionSystemIds `json:"system_ids"`

	// 获取密钥的DRM地址
	Url string `json:"url"`

	// drm speke 版本号 当前只支持1.0
	SpekeVersion EncryptionSpekeVersion `json:"speke_version"`

	// 请求模式。  取值如下： * direct_http：HTTP(S)直接访问DRM。 * functiongraph_proxy：FunctionGraph代理访问DRM。
	RequestMode EncryptionRequestMode `json:"request_mode"`

	// 需要添加在drm请求头中的鉴权信息。最多支持配置5个。  仅direct_http请求模式支持配置http_headers。
	HttpHeaders *[]HttpHeader `json:"http_headers,omitempty"`

	// functiongraph_proxy请求模式需要提供functiongraph的urn。
	Urn *string `json:"urn,omitempty"`
}

func (o Encryption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Encryption struct{}"
	}

	return strings.Join([]string{"Encryption", string(data)}, " ")
}

type EncryptionEncryptionMethod struct {
	value string
}

type EncryptionEncryptionMethodEnum struct {
	SAMPLE_AES EncryptionEncryptionMethod
	AES_128    EncryptionEncryptionMethod
}

func GetEncryptionEncryptionMethodEnum() EncryptionEncryptionMethodEnum {
	return EncryptionEncryptionMethodEnum{
		SAMPLE_AES: EncryptionEncryptionMethod{
			value: "SAMPLE-AES",
		},
		AES_128: EncryptionEncryptionMethod{
			value: "AES-128",
		},
	}
}

func (c EncryptionEncryptionMethod) Value() string {
	return c.value
}

func (c EncryptionEncryptionMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptionEncryptionMethod) UnmarshalJSON(b []byte) error {
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

type EncryptionLevel struct {
	value string
}

type EncryptionLevelEnum struct {
	CONTENT EncryptionLevel
	PROFILE EncryptionLevel
}

func GetEncryptionLevelEnum() EncryptionLevelEnum {
	return EncryptionLevelEnum{
		CONTENT: EncryptionLevel{
			value: "content",
		},
		PROFILE: EncryptionLevel{
			value: "profile",
		},
	}
}

func (c EncryptionLevel) Value() string {
	return c.value
}

func (c EncryptionLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptionLevel) UnmarshalJSON(b []byte) error {
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

type EncryptionSystemIds struct {
	value string
}

type EncryptionSystemIdsEnum struct {
	WIDEVINE   EncryptionSystemIds
	PLAY_READY EncryptionSystemIds
	FAIR_PLAY  EncryptionSystemIds
}

func GetEncryptionSystemIdsEnum() EncryptionSystemIdsEnum {
	return EncryptionSystemIdsEnum{
		WIDEVINE: EncryptionSystemIds{
			value: "Widevine",
		},
		PLAY_READY: EncryptionSystemIds{
			value: "PlayReady",
		},
		FAIR_PLAY: EncryptionSystemIds{
			value: "FairPlay",
		},
	}
}

func (c EncryptionSystemIds) Value() string {
	return c.value
}

func (c EncryptionSystemIds) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptionSystemIds) UnmarshalJSON(b []byte) error {
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

type EncryptionSpekeVersion struct {
	value string
}

type EncryptionSpekeVersionEnum struct {
	E_1_0 EncryptionSpekeVersion
}

func GetEncryptionSpekeVersionEnum() EncryptionSpekeVersionEnum {
	return EncryptionSpekeVersionEnum{
		E_1_0: EncryptionSpekeVersion{
			value: "1.0",
		},
	}
}

func (c EncryptionSpekeVersion) Value() string {
	return c.value
}

func (c EncryptionSpekeVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptionSpekeVersion) UnmarshalJSON(b []byte) error {
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

type EncryptionRequestMode struct {
	value string
}

type EncryptionRequestModeEnum struct {
	DIRECT_HTTP         EncryptionRequestMode
	FUNCTIONGRAPH_PROXY EncryptionRequestMode
}

func GetEncryptionRequestModeEnum() EncryptionRequestModeEnum {
	return EncryptionRequestModeEnum{
		DIRECT_HTTP: EncryptionRequestMode{
			value: "direct_http",
		},
		FUNCTIONGRAPH_PROXY: EncryptionRequestMode{
			value: "functiongraph_proxy",
		},
	}
}

func (c EncryptionRequestMode) Value() string {
	return c.value
}

func (c EncryptionRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptionRequestMode) UnmarshalJSON(b []byte) error {
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
