package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessWhiteList 进出登记白名单
type AccessWhiteList struct {

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 原因
	Reason *AccessWhiteListReason `json:"reason,omitempty"`

	// 原因补充
	ReasonSupplement *string `json:"reason_supplement,omitempty"`

	// 作业范围
	Scope *AccessWhiteListScope `json:"scope,omitempty"`

	// 机房编码
	RoomCode *string `json:"room_code,omitempty"`

	// 人员名单
	Whitelist *[]ContactInformation `json:"whitelist,omitempty"`
}

func (o AccessWhiteList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessWhiteList struct{}"
	}

	return strings.Join([]string{"AccessWhiteList", string(data)}, " ")
}

type AccessWhiteListReason struct {
	value string
}

type AccessWhiteListReasonEnum struct {
	E_1 AccessWhiteListReason
	E_2 AccessWhiteListReason
	E_3 AccessWhiteListReason
	E_4 AccessWhiteListReason
	E_5 AccessWhiteListReason
}

func GetAccessWhiteListReasonEnum() AccessWhiteListReasonEnum {
	return AccessWhiteListReasonEnum{
		E_1: AccessWhiteListReason{
			value: "1",
		},
		E_2: AccessWhiteListReason{
			value: "2",
		},
		E_3: AccessWhiteListReason{
			value: "3",
		},
		E_4: AccessWhiteListReason{
			value: "4",
		},
		E_5: AccessWhiteListReason{
			value: "5",
		},
	}
}

func (c AccessWhiteListReason) Value() string {
	return c.value
}

func (c AccessWhiteListReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessWhiteListReason) UnmarshalJSON(b []byte) error {
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

type AccessWhiteListScope struct {
	value string
}

type AccessWhiteListScopeEnum struct {
	ROOM   AccessWhiteListScope
	RACK   AccessWhiteListScope
	DEVICE AccessWhiteListScope
}

func GetAccessWhiteListScopeEnum() AccessWhiteListScopeEnum {
	return AccessWhiteListScopeEnum{
		ROOM: AccessWhiteListScope{
			value: "room",
		},
		RACK: AccessWhiteListScope{
			value: "rack",
		},
		DEVICE: AccessWhiteListScope{
			value: "device",
		},
	}
}

func (c AccessWhiteListScope) Value() string {
	return c.value
}

func (c AccessWhiteListScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessWhiteListScope) UnmarshalJSON(b []byte) error {
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
