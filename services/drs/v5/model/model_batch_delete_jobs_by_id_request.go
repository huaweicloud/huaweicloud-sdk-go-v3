package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteJobsByIdRequest Request Object
type BatchDeleteJobsByIdRequest struct {

	// 请求语言类型。
	XLanguage *BatchDeleteJobsByIdRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchDeleteJobReq `json:"body,omitempty"`
}

func (o BatchDeleteJobsByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsByIdRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsByIdRequest", string(data)}, " ")
}

type BatchDeleteJobsByIdRequestXLanguage struct {
	value string
}

type BatchDeleteJobsByIdRequestXLanguageEnum struct {
	EN_US BatchDeleteJobsByIdRequestXLanguage
	ZH_CN BatchDeleteJobsByIdRequestXLanguage
}

func GetBatchDeleteJobsByIdRequestXLanguageEnum() BatchDeleteJobsByIdRequestXLanguageEnum {
	return BatchDeleteJobsByIdRequestXLanguageEnum{
		EN_US: BatchDeleteJobsByIdRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchDeleteJobsByIdRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchDeleteJobsByIdRequestXLanguage) Value() string {
	return c.value
}

func (c BatchDeleteJobsByIdRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteJobsByIdRequestXLanguage) UnmarshalJSON(b []byte) error {
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
