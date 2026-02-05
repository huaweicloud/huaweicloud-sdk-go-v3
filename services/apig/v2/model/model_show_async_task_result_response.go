package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAsyncTaskResultResponse Response Object
type ShowAsyncTaskResultResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 异步任务状态。 - waiting：异步任务等待处理。 - processing：异步任务处理中。 - succeed：异步任务执行完成。 - failed：异步任务执行失败。
	TaskStatus *ShowAsyncTaskResultResponseTaskStatus `json:"task_status,omitempty"`

	// 任务类型。 import_api为异步导入api，export_api为异步导出api。
	TaskType *ShowAsyncTaskResultResponseTaskType `json:"task_type,omitempty"`

	// 任务结果。string可转成json object。 当task_type为import_api时，字段包括API分组编号group_id、success数组、failure数组、swagger结构体、ignore数组。   - success数组：导入成功的API，数组元素中包括4个字段，导入成功的API编号id、API请求方法method、API请求路径path、导入行为action（枚举值，update表示更新API，create表示新建API）。   - failure数组：导入失败的API，数组元素中包括4个字段，API请求方法method、API请求路径path、导入失败的错误码error_code、导入失败的错误信息error_msg。   - swagger结构体：包括2个字段，swagger文档编号id、导入结果说明result。   - ignore数组：元素包括API请求方法method、API请求路径path。    当task_type为export_api时，字段包括导出文件类型file_type、导出文件内容content。
	TaskResult     *string `json:"task_result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAsyncTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAsyncTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowAsyncTaskResultResponse", string(data)}, " ")
}

type ShowAsyncTaskResultResponseTaskStatus struct {
	value string
}

type ShowAsyncTaskResultResponseTaskStatusEnum struct {
	WAITING    ShowAsyncTaskResultResponseTaskStatus
	PROCESSING ShowAsyncTaskResultResponseTaskStatus
	SUCCEED    ShowAsyncTaskResultResponseTaskStatus
	FAILED     ShowAsyncTaskResultResponseTaskStatus
}

func GetShowAsyncTaskResultResponseTaskStatusEnum() ShowAsyncTaskResultResponseTaskStatusEnum {
	return ShowAsyncTaskResultResponseTaskStatusEnum{
		WAITING: ShowAsyncTaskResultResponseTaskStatus{
			value: "waiting",
		},
		PROCESSING: ShowAsyncTaskResultResponseTaskStatus{
			value: "processing",
		},
		SUCCEED: ShowAsyncTaskResultResponseTaskStatus{
			value: "succeed",
		},
		FAILED: ShowAsyncTaskResultResponseTaskStatus{
			value: "failed",
		},
	}
}

func (c ShowAsyncTaskResultResponseTaskStatus) Value() string {
	return c.value
}

func (c ShowAsyncTaskResultResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAsyncTaskResultResponseTaskStatus) UnmarshalJSON(b []byte) error {
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

type ShowAsyncTaskResultResponseTaskType struct {
	value string
}

type ShowAsyncTaskResultResponseTaskTypeEnum struct {
	IMPORT_API ShowAsyncTaskResultResponseTaskType
	EXPORT_API ShowAsyncTaskResultResponseTaskType
}

func GetShowAsyncTaskResultResponseTaskTypeEnum() ShowAsyncTaskResultResponseTaskTypeEnum {
	return ShowAsyncTaskResultResponseTaskTypeEnum{
		IMPORT_API: ShowAsyncTaskResultResponseTaskType{
			value: "import_api",
		},
		EXPORT_API: ShowAsyncTaskResultResponseTaskType{
			value: "export_api",
		},
	}
}

func (c ShowAsyncTaskResultResponseTaskType) Value() string {
	return c.value
}

func (c ShowAsyncTaskResultResponseTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAsyncTaskResultResponseTaskType) UnmarshalJSON(b []byte) error {
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
