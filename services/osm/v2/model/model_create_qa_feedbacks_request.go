package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateQaFeedbacksRequest struct {

	// 调用智能客服服务标志。
	XServiceKey *string `json:"x-service-key,omitempty"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// - LIKE:  - CANCEL_LIKE:  - DISLIKE:  - ALL:
	FeedbackType CreateQaFeedbacksRequestFeedbackType `json:"feedback_type"`

	Body *QaFeedbackReq `json:"body,omitempty"`
}

func (o CreateQaFeedbacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQaFeedbacksRequest struct{}"
	}

	return strings.Join([]string{"CreateQaFeedbacksRequest", string(data)}, " ")
}

type CreateQaFeedbacksRequestFeedbackType struct {
	value string
}

type CreateQaFeedbacksRequestFeedbackTypeEnum struct {
	LIKE        CreateQaFeedbacksRequestFeedbackType
	CANCEL_LIKE CreateQaFeedbacksRequestFeedbackType
	DISLIKE     CreateQaFeedbacksRequestFeedbackType
	ALL         CreateQaFeedbacksRequestFeedbackType
}

func GetCreateQaFeedbacksRequestFeedbackTypeEnum() CreateQaFeedbacksRequestFeedbackTypeEnum {
	return CreateQaFeedbacksRequestFeedbackTypeEnum{
		LIKE: CreateQaFeedbacksRequestFeedbackType{
			value: "LIKE",
		},
		CANCEL_LIKE: CreateQaFeedbacksRequestFeedbackType{
			value: "CANCEL_LIKE",
		},
		DISLIKE: CreateQaFeedbacksRequestFeedbackType{
			value: "DISLIKE",
		},
		ALL: CreateQaFeedbacksRequestFeedbackType{
			value: "ALL",
		},
	}
}

func (c CreateQaFeedbacksRequestFeedbackType) Value() string {
	return c.value
}

func (c CreateQaFeedbacksRequestFeedbackType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateQaFeedbacksRequestFeedbackType) UnmarshalJSON(b []byte) error {
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
