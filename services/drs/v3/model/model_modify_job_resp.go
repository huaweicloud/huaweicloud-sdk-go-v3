package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type ModifyJobResp struct {
	// 任务ID

	Id string `json:"id"`
	// 状态

	Status ModifyJobRespStatus `json:"status"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ModifyJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobResp struct{}"
	}

	return strings.Join([]string{"ModifyJobResp", string(data)}, " ")
}

type ModifyJobRespStatus struct {
	value string
}

type ModifyJobRespStatusEnum struct {
	SUCCESS ModifyJobRespStatus
	FAILED  ModifyJobRespStatus
}

func GetModifyJobRespStatusEnum() ModifyJobRespStatusEnum {
	return ModifyJobRespStatusEnum{
		SUCCESS: ModifyJobRespStatus{
			value: "success 成功",
		},
		FAILED: ModifyJobRespStatus{
			value: "failed 失败",
		},
	}
}

func (c ModifyJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobRespStatus) UnmarshalJSON(b []byte) error {
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
