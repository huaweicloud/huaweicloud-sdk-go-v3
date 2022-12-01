package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
