package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SynchronizeInstancesRequest Request Object
type SynchronizeInstancesRequest struct {

	// 请求语言类型。
	XLanguage *SynchronizeInstancesRequestXLanguage `json:"X-Language,omitempty"`

	Body *SynchronizeInstancesReq `json:"body,omitempty"`
}

func (o SynchronizeInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynchronizeInstancesRequest struct{}"
	}

	return strings.Join([]string{"SynchronizeInstancesRequest", string(data)}, " ")
}

type SynchronizeInstancesRequestXLanguage struct {
	value string
}

type SynchronizeInstancesRequestXLanguageEnum struct {
	EN_US SynchronizeInstancesRequestXLanguage
	ZH_CN SynchronizeInstancesRequestXLanguage
}

func GetSynchronizeInstancesRequestXLanguageEnum() SynchronizeInstancesRequestXLanguageEnum {
	return SynchronizeInstancesRequestXLanguageEnum{
		EN_US: SynchronizeInstancesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: SynchronizeInstancesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c SynchronizeInstancesRequestXLanguage) Value() string {
	return c.value
}

func (c SynchronizeInstancesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SynchronizeInstancesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
