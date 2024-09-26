package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateHarvestJobStatusRequestBody 修改状态请求体
type UpdateHarvestJobStatusRequestBody struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 任务状态
	Status UpdateHarvestJobStatusRequestBodyStatus `json:"status"`
}

func (o UpdateHarvestJobStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHarvestJobStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHarvestJobStatusRequestBody", string(data)}, " ")
}

type UpdateHarvestJobStatusRequestBodyStatus struct {
	value string
}

type UpdateHarvestJobStatusRequestBodyStatusEnum struct {
	UNSTART UpdateHarvestJobStatusRequestBodyStatus
}

func GetUpdateHarvestJobStatusRequestBodyStatusEnum() UpdateHarvestJobStatusRequestBodyStatusEnum {
	return UpdateHarvestJobStatusRequestBodyStatusEnum{
		UNSTART: UpdateHarvestJobStatusRequestBodyStatus{
			value: "UNSTART",
		},
	}
}

func (c UpdateHarvestJobStatusRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateHarvestJobStatusRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHarvestJobStatusRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
