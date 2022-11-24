package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchUpdateUserStatusRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 用户的新状态。 - 0，启用 - 1，停用
	Value BatchUpdateUserStatusRequestValue `json:"value"`

	// 帐号类型。默认0。 * 0：华为云会议帐号。用于帐号/密码鉴权方式 * 1：第三方User ID，用于App ID鉴权方式
	AccountType *int32 `json:"accountType,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchUpdateUserStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUserStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateUserStatusRequest", string(data)}, " ")
}

type BatchUpdateUserStatusRequestValue struct {
	value int32
}

type BatchUpdateUserStatusRequestValueEnum struct {
	E_0 BatchUpdateUserStatusRequestValue
	E_1 BatchUpdateUserStatusRequestValue
}

func GetBatchUpdateUserStatusRequestValueEnum() BatchUpdateUserStatusRequestValueEnum {
	return BatchUpdateUserStatusRequestValueEnum{
		E_0: BatchUpdateUserStatusRequestValue{
			value: 0,
		}, E_1: BatchUpdateUserStatusRequestValue{
			value: 1,
		},
	}
}

func (c BatchUpdateUserStatusRequestValue) Value() int32 {
	return c.value
}

func (c BatchUpdateUserStatusRequestValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateUserStatusRequestValue) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
