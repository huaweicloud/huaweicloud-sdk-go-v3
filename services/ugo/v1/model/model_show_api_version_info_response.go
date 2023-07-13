package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiVersionInfoResponse Response Object
type ShowApiVersionInfoResponse struct {

	// 版本ID。
	Id *string `json:"id,omitempty"`

	// 版本详情的URL地址。
	Links *string `json:"links,omitempty"`

	// 该版本API的微版本信息。
	Version *string `json:"version,omitempty"`

	// 版本的状态。
	Status *ShowApiVersionInfoResponseStatus `json:"status,omitempty"`

	// 版本更新时间。
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowApiVersionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionInfoResponse", string(data)}, " ")
}

type ShowApiVersionInfoResponseStatus struct {
	value string
}

type ShowApiVersionInfoResponseStatusEnum struct {
	CURRENT    ShowApiVersionInfoResponseStatus
	SUPPORTED  ShowApiVersionInfoResponseStatus
	DEPRECATED ShowApiVersionInfoResponseStatus
}

func GetShowApiVersionInfoResponseStatusEnum() ShowApiVersionInfoResponseStatusEnum {
	return ShowApiVersionInfoResponseStatusEnum{
		CURRENT: ShowApiVersionInfoResponseStatus{
			value: "CURRENT",
		},
		SUPPORTED: ShowApiVersionInfoResponseStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: ShowApiVersionInfoResponseStatus{
			value: "DEPRECATED",
		},
	}
}

func (c ShowApiVersionInfoResponseStatus) Value() string {
	return c.value
}

func (c ShowApiVersionInfoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiVersionInfoResponseStatus) UnmarshalJSON(b []byte) error {
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
