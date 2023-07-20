package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloneJobReq 克隆任务请求体。
type CloneJobReq struct {

	// 被克隆任务ID。
	JobId string `json:"job_id"`

	// 克隆任务名称。名称在4位到50位之间，必须以字母开头，可以包含字母、数字、中划线或下划线，不能包含其他特殊字符，任务名称不能重复。
	Name string `json:"name"`

	// 任务版本号，新UX任务为2.0。默认为空，即克隆老任务。
	TaskVersion *CloneJobReqTaskVersion `json:"task_version,omitempty"`
}

func (o CloneJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneJobReq struct{}"
	}

	return strings.Join([]string{"CloneJobReq", string(data)}, " ")
}

type CloneJobReqTaskVersion struct {
	value string
}

type CloneJobReqTaskVersionEnum struct {
	E_2_0 CloneJobReqTaskVersion
}

func GetCloneJobReqTaskVersionEnum() CloneJobReqTaskVersionEnum {
	return CloneJobReqTaskVersionEnum{
		E_2_0: CloneJobReqTaskVersion{
			value: "2.0",
		},
	}
}

func (c CloneJobReqTaskVersion) Value() string {
	return c.value
}

func (c CloneJobReqTaskVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloneJobReqTaskVersion) UnmarshalJSON(b []byte) error {
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
