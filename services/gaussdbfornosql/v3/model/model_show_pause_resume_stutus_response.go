package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowPauseResumeStutusResponse struct {

	// 容灾实例数据同步状态 - NA：实例尚未搭建容灾关系 - NEW：尚未启动的数据同步状态 - SYNCING：数据同步正常进行中 - SUSPENDING：正在暂停数据同步 - SUSPENDED：数据同步已暂停 - RECOVERYING：正在恢复数据同步
	Status         *ShowPauseResumeStutusResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowPauseResumeStutusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPauseResumeStutusResponse struct{}"
	}

	return strings.Join([]string{"ShowPauseResumeStutusResponse", string(data)}, " ")
}

type ShowPauseResumeStutusResponseStatus struct {
	value string
}

type ShowPauseResumeStutusResponseStatusEnum struct {
	NA          ShowPauseResumeStutusResponseStatus
	NEW         ShowPauseResumeStutusResponseStatus
	SYNCING     ShowPauseResumeStutusResponseStatus
	SUSPENDING  ShowPauseResumeStutusResponseStatus
	SUSPENDED   ShowPauseResumeStutusResponseStatus
	RECOVERYING ShowPauseResumeStutusResponseStatus
}

func GetShowPauseResumeStutusResponseStatusEnum() ShowPauseResumeStutusResponseStatusEnum {
	return ShowPauseResumeStutusResponseStatusEnum{
		NA: ShowPauseResumeStutusResponseStatus{
			value: "NA",
		},
		NEW: ShowPauseResumeStutusResponseStatus{
			value: "NEW",
		},
		SYNCING: ShowPauseResumeStutusResponseStatus{
			value: "SYNCING",
		},
		SUSPENDING: ShowPauseResumeStutusResponseStatus{
			value: "SUSPENDING",
		},
		SUSPENDED: ShowPauseResumeStutusResponseStatus{
			value: "SUSPENDED",
		},
		RECOVERYING: ShowPauseResumeStutusResponseStatus{
			value: "RECOVERYING",
		},
	}
}

func (c ShowPauseResumeStutusResponseStatus) Value() string {
	return c.value
}

func (c ShowPauseResumeStutusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPauseResumeStutusResponseStatus) UnmarshalJSON(b []byte) error {
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
