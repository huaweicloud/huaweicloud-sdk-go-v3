package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHpcCacheTaskResponse Response Object
type ShowHpcCacheTaskResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型
	Type *string `json:"type,omitempty"`

	// 任务状态
	Status *ShowHpcCacheTaskResponseStatus `json:"status,omitempty"`

	// 联动目录名称
	SrcTarget *string `json:"src_target,omitempty"`

	// 导入导出任务的源端路径前缀
	SrcPrefix *string `json:"src_prefix,omitempty"`

	// 和src_target保持一致
	DestTarget *string `json:"dest_target,omitempty"`

	// 和src_prefix保持一致
	DestPrefix *string `json:"dest_prefix,omitempty"`

	// 任务开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务执行结果信息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHpcCacheTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHpcCacheTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowHpcCacheTaskResponse", string(data)}, " ")
}

type ShowHpcCacheTaskResponseStatus struct {
	value string
}

type ShowHpcCacheTaskResponseStatusEnum struct {
	SUCCESS ShowHpcCacheTaskResponseStatus
	DOING   ShowHpcCacheTaskResponseStatus
	FAIL    ShowHpcCacheTaskResponseStatus
}

func GetShowHpcCacheTaskResponseStatusEnum() ShowHpcCacheTaskResponseStatusEnum {
	return ShowHpcCacheTaskResponseStatusEnum{
		SUCCESS: ShowHpcCacheTaskResponseStatus{
			value: "SUCCESS",
		},
		DOING: ShowHpcCacheTaskResponseStatus{
			value: "DOING",
		},
		FAIL: ShowHpcCacheTaskResponseStatus{
			value: "FAIL",
		},
	}
}

func (c ShowHpcCacheTaskResponseStatus) Value() string {
	return c.value
}

func (c ShowHpcCacheTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHpcCacheTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
