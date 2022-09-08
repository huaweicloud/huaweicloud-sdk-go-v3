package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowJobStatusResponse struct {
	Name *string `json:"name,omitempty"`

	Status *ShowJobStatusResponseStatus `json:"status,omitempty"`

	Starttime *string `json:"starttime,omitempty"`

	EndTime *string `json:"endTime,omitempty"`

	// 状态最后更新时间
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Nodes          *[]RealTimeNodeStatus `json:"nodes,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}

type ShowJobStatusResponseStatus struct {
	value string
}

type ShowJobStatusResponseStatusEnum struct {
	STARTTING ShowJobStatusResponseStatus
	NORMAL    ShowJobStatusResponseStatus
	EXCEPTION ShowJobStatusResponseStatus
	STOPPING  ShowJobStatusResponseStatus
	STOPPED   ShowJobStatusResponseStatus
}

func GetShowJobStatusResponseStatusEnum() ShowJobStatusResponseStatusEnum {
	return ShowJobStatusResponseStatusEnum{
		STARTTING: ShowJobStatusResponseStatus{
			value: "STARTTING",
		},
		NORMAL: ShowJobStatusResponseStatus{
			value: "NORMAL",
		},
		EXCEPTION: ShowJobStatusResponseStatus{
			value: "EXCEPTION",
		},
		STOPPING: ShowJobStatusResponseStatus{
			value: "STOPPING",
		},
		STOPPED: ShowJobStatusResponseStatus{
			value: "STOPPED",
		},
	}
}

func (c ShowJobStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
