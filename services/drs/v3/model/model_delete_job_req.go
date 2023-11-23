package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteJobReq 删除在线迁移任务请求体
type DeleteJobReq struct {

	// terminate:结束迁移任务,force_terminate:强制结束迁移任务,delete:删除迁移任务
	DeleteType DeleteJobReqDeleteType `json:"delete_type"`

	// 任务ID
	JobId string `json:"job_id"`

	// MySQL为源，实时迁移，实时同步，数据订阅，实时灾备结束任务时是否展示断点信息
	IsShowBreakpointPosition *bool `json:"is_show_breakpoint_position,omitempty"`
}

func (o DeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobReq struct{}"
	}

	return strings.Join([]string{"DeleteJobReq", string(data)}, " ")
}

type DeleteJobReqDeleteType struct {
	value string
}

type DeleteJobReqDeleteTypeEnum struct {
	TERMINATE       DeleteJobReqDeleteType
	FORCE_TERMINATE DeleteJobReqDeleteType
	DELETE          DeleteJobReqDeleteType
}

func GetDeleteJobReqDeleteTypeEnum() DeleteJobReqDeleteTypeEnum {
	return DeleteJobReqDeleteTypeEnum{
		TERMINATE: DeleteJobReqDeleteType{
			value: "terminate",
		},
		FORCE_TERMINATE: DeleteJobReqDeleteType{
			value: "force_terminate",
		},
		DELETE: DeleteJobReqDeleteType{
			value: "delete",
		},
	}
}

func (c DeleteJobReqDeleteType) Value() string {
	return c.value
}

func (c DeleteJobReqDeleteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobReqDeleteType) UnmarshalJSON(b []byte) error {
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
