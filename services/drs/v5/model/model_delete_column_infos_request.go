package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteColumnInfosRequest Request Object
type DeleteColumnInfosRequest struct {

	// 请求语言类型。
	XLanguage *DeleteColumnInfosRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *DeleteColumnInfoReq `json:"body,omitempty"`
}

func (o DeleteColumnInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteColumnInfosRequest struct{}"
	}

	return strings.Join([]string{"DeleteColumnInfosRequest", string(data)}, " ")
}

type DeleteColumnInfosRequestXLanguage struct {
	value string
}

type DeleteColumnInfosRequestXLanguageEnum struct {
	EN_US DeleteColumnInfosRequestXLanguage
	ZH_CN DeleteColumnInfosRequestXLanguage
}

func GetDeleteColumnInfosRequestXLanguageEnum() DeleteColumnInfosRequestXLanguageEnum {
	return DeleteColumnInfosRequestXLanguageEnum{
		EN_US: DeleteColumnInfosRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteColumnInfosRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteColumnInfosRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteColumnInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteColumnInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
