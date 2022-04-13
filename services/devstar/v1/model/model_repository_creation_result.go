package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 仓库信息
type RepositoryCreationResult struct {
	Repository *RepositoryBasicInfo `json:"repository,omitempty"`
	// 任务id

	TaskId *string `json:"task_id,omitempty"`
	// 任务状态, success:成功,failed:失败,creating:创建中

	Status *RepositoryCreationResultStatus `json:"status,omitempty"`
	// 失败原因

	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o RepositoryCreationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryCreationResult struct{}"
	}

	return strings.Join([]string{"RepositoryCreationResult", string(data)}, " ")
}

type RepositoryCreationResultStatus struct {
	value string
}

type RepositoryCreationResultStatusEnum struct {
	SUCCESS  RepositoryCreationResultStatus
	FAILED   RepositoryCreationResultStatus
	CREATING RepositoryCreationResultStatus
}

func GetRepositoryCreationResultStatusEnum() RepositoryCreationResultStatusEnum {
	return RepositoryCreationResultStatusEnum{
		SUCCESS: RepositoryCreationResultStatus{
			value: "success",
		},
		FAILED: RepositoryCreationResultStatus{
			value: "failed",
		},
		CREATING: RepositoryCreationResultStatus{
			value: "creating",
		},
	}
}

func (c RepositoryCreationResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryCreationResultStatus) UnmarshalJSON(b []byte) error {
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
