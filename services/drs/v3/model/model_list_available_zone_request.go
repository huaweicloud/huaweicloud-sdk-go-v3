package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAvailableZoneRequest Request Object
type ListAvailableZoneRequest struct {

	// 请求语言类型。
	XLanguage *ListAvailableZoneRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryAvailableNodeTypeReq `json:"body,omitempty"`
}

func (o ListAvailableZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZoneRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableZoneRequest", string(data)}, " ")
}

type ListAvailableZoneRequestXLanguage struct {
	value string
}

type ListAvailableZoneRequestXLanguageEnum struct {
	EN_US ListAvailableZoneRequestXLanguage
	ZH_CN ListAvailableZoneRequestXLanguage
}

func GetListAvailableZoneRequestXLanguageEnum() ListAvailableZoneRequestXLanguageEnum {
	return ListAvailableZoneRequestXLanguageEnum{
		EN_US: ListAvailableZoneRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListAvailableZoneRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListAvailableZoneRequestXLanguage) Value() string {
	return c.value
}

func (c ListAvailableZoneRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAvailableZoneRequestXLanguage) UnmarshalJSON(b []byte) error {
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
