package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProvidersRequest Request Object
type ListProvidersRequest struct {

	// 分页偏移
	Offset *int32 `json:"offset,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 资源是否默认收集
	Track *ListProvidersRequestTrack `json:"track,omitempty"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *ListProvidersRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListProvidersRequest", string(data)}, " ")
}

type ListProvidersRequestTrack struct {
	value string
}

type ListProvidersRequestTrackEnum struct {
	TRACKED   ListProvidersRequestTrack
	UNTRACKED ListProvidersRequestTrack
}

func GetListProvidersRequestTrackEnum() ListProvidersRequestTrackEnum {
	return ListProvidersRequestTrackEnum{
		TRACKED: ListProvidersRequestTrack{
			value: "tracked",
		},
		UNTRACKED: ListProvidersRequestTrack{
			value: "untracked",
		},
	}
}

func (c ListProvidersRequestTrack) Value() string {
	return c.value
}

func (c ListProvidersRequestTrack) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProvidersRequestTrack) UnmarshalJSON(b []byte) error {
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

type ListProvidersRequestXLanguage struct {
	value string
}

type ListProvidersRequestXLanguageEnum struct {
	ZH_CN ListProvidersRequestXLanguage
	EN_US ListProvidersRequestXLanguage
	FR_FR ListProvidersRequestXLanguage
	ES_US ListProvidersRequestXLanguage
	PT_BR ListProvidersRequestXLanguage
}

func GetListProvidersRequestXLanguageEnum() ListProvidersRequestXLanguageEnum {
	return ListProvidersRequestXLanguageEnum{
		ZH_CN: ListProvidersRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListProvidersRequestXLanguage{
			value: "en-us",
		},
		FR_FR: ListProvidersRequestXLanguage{
			value: "fr-fr",
		},
		ES_US: ListProvidersRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ListProvidersRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ListProvidersRequestXLanguage) Value() string {
	return c.value
}

func (c ListProvidersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProvidersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
