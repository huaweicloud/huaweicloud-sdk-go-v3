package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTagStatusResponse Response Object
type ShowTagStatusResponse struct {

	// 标签处理状态
	Status *ShowTagStatusResponseStatus `json:"status,omitempty"`

	// 默认标签是否已开启
	DefaultTagsEnable *bool `json:"default_tags_enable,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowTagStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowTagStatusResponse", string(data)}, " ")
}

type ShowTagStatusResponseStatus struct {
	value string
}

type ShowTagStatusResponseStatusEnum struct {
	PROCESSING ShowTagStatusResponseStatus
	SUCCEED    ShowTagStatusResponseStatus
	FAILED     ShowTagStatusResponseStatus
}

func GetShowTagStatusResponseStatusEnum() ShowTagStatusResponseStatusEnum {
	return ShowTagStatusResponseStatusEnum{
		PROCESSING: ShowTagStatusResponseStatus{
			value: "processing",
		},
		SUCCEED: ShowTagStatusResponseStatus{
			value: "succeed",
		},
		FAILED: ShowTagStatusResponseStatus{
			value: "failed",
		},
	}
}

func (c ShowTagStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowTagStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTagStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
