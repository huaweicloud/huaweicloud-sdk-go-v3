package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUpdateObjectSavingStatusResponse Response Object
type ShowUpdateObjectSavingStatusResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态 pending：处理中 failed：失败 success：成功
	Status *ShowUpdateObjectSavingStatusResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUpdateObjectSavingStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateObjectSavingStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowUpdateObjectSavingStatusResponse", string(data)}, " ")
}

type ShowUpdateObjectSavingStatusResponseStatus struct {
	value string
}

type ShowUpdateObjectSavingStatusResponseStatusEnum struct {
	PENDING ShowUpdateObjectSavingStatusResponseStatus
	FAILED  ShowUpdateObjectSavingStatusResponseStatus
	SUCCESS ShowUpdateObjectSavingStatusResponseStatus
}

func GetShowUpdateObjectSavingStatusResponseStatusEnum() ShowUpdateObjectSavingStatusResponseStatusEnum {
	return ShowUpdateObjectSavingStatusResponseStatusEnum{
		PENDING: ShowUpdateObjectSavingStatusResponseStatus{
			value: "pending",
		},
		FAILED: ShowUpdateObjectSavingStatusResponseStatus{
			value: "failed",
		},
		SUCCESS: ShowUpdateObjectSavingStatusResponseStatus{
			value: "success",
		},
	}
}

func (c ShowUpdateObjectSavingStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowUpdateObjectSavingStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUpdateObjectSavingStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
