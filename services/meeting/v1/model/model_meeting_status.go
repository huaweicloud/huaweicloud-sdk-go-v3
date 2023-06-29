package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MeetingStatus 会议状态。 - schedule:预定状态 - created:会议已经被创建并正在召开 - destroyed:会议已经关闭
type MeetingStatus struct {
	value string
}

type MeetingStatusEnum struct {
	SCHEDULE  MeetingStatus
	CREATED   MeetingStatus
	DESTROYED MeetingStatus
}

func GetMeetingStatusEnum() MeetingStatusEnum {
	return MeetingStatusEnum{
		SCHEDULE: MeetingStatus{
			value: "schedule",
		},
		CREATED: MeetingStatus{
			value: "created",
		},
		DESTROYED: MeetingStatus{
			value: "destroyed",
		},
	}
}

func (c MeetingStatus) Value() string {
	return c.value
}

func (c MeetingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MeetingStatus) UnmarshalJSON(b []byte) error {
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
