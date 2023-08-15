package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AsyncJobResp 批量异步创建任务响应体。
type AsyncJobResp struct {

	// 批量异步创建的任务ID。
	AsyncJobId string `json:"async_job_id"`

	// 批量异步创建的任务状态。取值： - ASYNC_JOB_VALIDATING：批量异步任务参数校验中。 - ASYNC_JOB_VALIDATE_FAILED：批量异步任务参数校验失败。 - AUTO_PARAM_VALIDATE_SUCCESS：批量异步任务参数校验成功。 - COMMIT_SUCCESS：批量异步任务提交成功。
	Status AsyncJobRespStatus `json:"status"`

	// 批量异步创建的任务的租户名。
	DomainName string `json:"domain_name"`

	// 批量异步创建的任务的用户名。
	UserName string `json:"user_name"`

	// 批量异步创建的任务的创建时间。
	CreateTime string `json:"create_time"`
}

func (o AsyncJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncJobResp struct{}"
	}

	return strings.Join([]string{"AsyncJobResp", string(data)}, " ")
}

type AsyncJobRespStatus struct {
	value string
}

type AsyncJobRespStatusEnum struct {
	ASYNC_JOB_VALIDATING        AsyncJobRespStatus
	ASYNC_JOB_VALIDATE_FAILED   AsyncJobRespStatus
	AUTO_PARAM_VALIDATE_SUCCESS AsyncJobRespStatus
	COMMIT_SUCCESS              AsyncJobRespStatus
}

func GetAsyncJobRespStatusEnum() AsyncJobRespStatusEnum {
	return AsyncJobRespStatusEnum{
		ASYNC_JOB_VALIDATING: AsyncJobRespStatus{
			value: "ASYNC_JOB_VALIDATING",
		},
		ASYNC_JOB_VALIDATE_FAILED: AsyncJobRespStatus{
			value: "ASYNC_JOB_VALIDATE_FAILED",
		},
		AUTO_PARAM_VALIDATE_SUCCESS: AsyncJobRespStatus{
			value: "AUTO_PARAM_VALIDATE_SUCCESS",
		},
		COMMIT_SUCCESS: AsyncJobRespStatus{
			value: "COMMIT_SUCCESS",
		},
	}
}

func (c AsyncJobRespStatus) Value() string {
	return c.value
}

func (c AsyncJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AsyncJobRespStatus) UnmarshalJSON(b []byte) error {
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
