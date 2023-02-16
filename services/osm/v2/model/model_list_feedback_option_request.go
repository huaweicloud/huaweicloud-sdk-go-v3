package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFeedbackOptionRequest struct {

	// 调用智能客服服务标志。
	XServiceKey *string `json:"x-service-key,omitempty"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// - UNPUBLISHED:  - PUBLISH:
	Status ListFeedbackOptionRequestStatus `json:"status"`

	// - FAQ:  - FLOW:
	FeedbackSource *ListFeedbackOptionRequestFeedbackSource `json:"feedback_source,omitempty"`
}

func (o ListFeedbackOptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeedbackOptionRequest struct{}"
	}

	return strings.Join([]string{"ListFeedbackOptionRequest", string(data)}, " ")
}

type ListFeedbackOptionRequestStatus struct {
	value string
}

type ListFeedbackOptionRequestStatusEnum struct {
	UNPUBLISHED ListFeedbackOptionRequestStatus
	PUBLISH     ListFeedbackOptionRequestStatus
}

func GetListFeedbackOptionRequestStatusEnum() ListFeedbackOptionRequestStatusEnum {
	return ListFeedbackOptionRequestStatusEnum{
		UNPUBLISHED: ListFeedbackOptionRequestStatus{
			value: "UNPUBLISHED",
		},
		PUBLISH: ListFeedbackOptionRequestStatus{
			value: "PUBLISH",
		},
	}
}

func (c ListFeedbackOptionRequestStatus) Value() string {
	return c.value
}

func (c ListFeedbackOptionRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFeedbackOptionRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListFeedbackOptionRequestFeedbackSource struct {
	value string
}

type ListFeedbackOptionRequestFeedbackSourceEnum struct {
	FAQ  ListFeedbackOptionRequestFeedbackSource
	FLOW ListFeedbackOptionRequestFeedbackSource
}

func GetListFeedbackOptionRequestFeedbackSourceEnum() ListFeedbackOptionRequestFeedbackSourceEnum {
	return ListFeedbackOptionRequestFeedbackSourceEnum{
		FAQ: ListFeedbackOptionRequestFeedbackSource{
			value: "FAQ",
		},
		FLOW: ListFeedbackOptionRequestFeedbackSource{
			value: "FLOW",
		},
	}
}

func (c ListFeedbackOptionRequestFeedbackSource) Value() string {
	return c.value
}

func (c ListFeedbackOptionRequestFeedbackSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFeedbackOptionRequestFeedbackSource) UnmarshalJSON(b []byte) error {
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
