package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 会议状态。 - schedule:预定状态 - created:会议已经被创建并正在召开 - destroyed:会议已经关闭
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

func (c MeetingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MeetingStatus) UnmarshalJSON(b []byte) error {
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
