package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDetailsOfInstanceProgressV2Response struct {

	// 实例创建进度  单位：百分比
	Progress *ShowDetailsOfInstanceProgressV2ResponseProgress `json:"progress,omitempty" xml:"progress"`

	// 实例创建状态 - creating：创建中 - success：创建成功 - failed：创建失败
	Status *ShowDetailsOfInstanceProgressV2ResponseStatus `json:"status,omitempty" xml:"status"`

	// 实例创建失败错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 实例创建失败错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 实例创建开始时间。unix时间戳格式。
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 实例创建结束时间。unix时间戳格式。
	EndTime        *int64 `json:"end_time,omitempty" xml:"end_time"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDetailsOfInstanceProgressV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceProgressV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceProgressV2Response", string(data)}, " ")
}

type ShowDetailsOfInstanceProgressV2ResponseProgress struct {
	value int32
}

type ShowDetailsOfInstanceProgressV2ResponseProgressEnum struct {
	E_30  ShowDetailsOfInstanceProgressV2ResponseProgress
	E_50  ShowDetailsOfInstanceProgressV2ResponseProgress
	E_80  ShowDetailsOfInstanceProgressV2ResponseProgress
	E_90  ShowDetailsOfInstanceProgressV2ResponseProgress
	E_100 ShowDetailsOfInstanceProgressV2ResponseProgress
}

func GetShowDetailsOfInstanceProgressV2ResponseProgressEnum() ShowDetailsOfInstanceProgressV2ResponseProgressEnum {
	return ShowDetailsOfInstanceProgressV2ResponseProgressEnum{
		E_30: ShowDetailsOfInstanceProgressV2ResponseProgress{
			value: 30,
		}, E_50: ShowDetailsOfInstanceProgressV2ResponseProgress{
			value: 50,
		}, E_80: ShowDetailsOfInstanceProgressV2ResponseProgress{
			value: 80,
		}, E_90: ShowDetailsOfInstanceProgressV2ResponseProgress{
			value: 90,
		}, E_100: ShowDetailsOfInstanceProgressV2ResponseProgress{
			value: 100,
		},
	}
}

func (c ShowDetailsOfInstanceProgressV2ResponseProgress) Value() int32 {
	return c.value
}

func (c ShowDetailsOfInstanceProgressV2ResponseProgress) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceProgressV2ResponseProgress) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfInstanceProgressV2ResponseStatus struct {
	value string
}

type ShowDetailsOfInstanceProgressV2ResponseStatusEnum struct {
	CREATING ShowDetailsOfInstanceProgressV2ResponseStatus
	SUCCESS  ShowDetailsOfInstanceProgressV2ResponseStatus
	FAILED   ShowDetailsOfInstanceProgressV2ResponseStatus
}

func GetShowDetailsOfInstanceProgressV2ResponseStatusEnum() ShowDetailsOfInstanceProgressV2ResponseStatusEnum {
	return ShowDetailsOfInstanceProgressV2ResponseStatusEnum{
		CREATING: ShowDetailsOfInstanceProgressV2ResponseStatus{
			value: "creating",
		},
		SUCCESS: ShowDetailsOfInstanceProgressV2ResponseStatus{
			value: "success",
		},
		FAILED: ShowDetailsOfInstanceProgressV2ResponseStatus{
			value: "failed",
		},
	}
}

func (c ShowDetailsOfInstanceProgressV2ResponseStatus) Value() string {
	return c.value
}

func (c ShowDetailsOfInstanceProgressV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceProgressV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
