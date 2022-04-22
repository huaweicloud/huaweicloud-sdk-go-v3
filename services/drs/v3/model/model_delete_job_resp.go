package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 删除或者结束任务返回体
type DeleteJobResp struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 状态
	Status *DeleteJobRespStatus `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DeleteJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResp struct{}"
	}

	return strings.Join([]string{"DeleteJobResp", string(data)}, " ")
}

type DeleteJobRespStatus struct {
	value string
}

type DeleteJobRespStatusEnum struct {
	SUCCESS DeleteJobRespStatus
	FAILED  DeleteJobRespStatus
}

func GetDeleteJobRespStatusEnum() DeleteJobRespStatusEnum {
	return DeleteJobRespStatusEnum{
		SUCCESS: DeleteJobRespStatus{
			value: "success",
		},
		FAILED: DeleteJobRespStatus{
			value: "failed",
		},
	}
}

func (c DeleteJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobRespStatus) UnmarshalJSON(b []byte) error {
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
