package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListQuotasRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ListQuotasRequestContentType `json:"Content-Type"`
}

func (o ListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}

type ListQuotasRequestContentType struct {
	value string
}

type ListQuotasRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListQuotasRequestContentType
	APPLICATION_JSON             ListQuotasRequestContentType
}

func GetListQuotasRequestContentTypeEnum() ListQuotasRequestContentTypeEnum {
	return ListQuotasRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListQuotasRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListQuotasRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListQuotasRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQuotasRequestContentType) UnmarshalJSON(b []byte) error {
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
