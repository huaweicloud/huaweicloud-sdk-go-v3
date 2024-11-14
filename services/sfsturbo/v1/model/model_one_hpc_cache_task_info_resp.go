package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OneHpcCacheTaskInfoResp 每个任务信息
type OneHpcCacheTaskInfoResp struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 任务类型
	Type string `json:"type"`

	// 任务状态
	Status OneHpcCacheTaskInfoRespStatus `json:"status"`

	// 联动目录名称
	SrcTarget string `json:"src_target"`

	// 导入导出任务的源端路径前缀
	SrcPrefix string `json:"src_prefix"`

	// 和src_target保持一致
	DestTarget string `json:"dest_target"`

	// 和src_prefix保持一致
	DestPrefix string `json:"dest_prefix"`

	// 任务开始时间
	StartTime string `json:"start_time"`

	// 任务结束时间
	EndTime string `json:"end_time"`

	// 任务执行结果信息
	Message string `json:"message"`
}

func (o OneHpcCacheTaskInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneHpcCacheTaskInfoResp struct{}"
	}

	return strings.Join([]string{"OneHpcCacheTaskInfoResp", string(data)}, " ")
}

type OneHpcCacheTaskInfoRespStatus struct {
	value string
}

type OneHpcCacheTaskInfoRespStatusEnum struct {
	SUCCESS OneHpcCacheTaskInfoRespStatus
	DOING   OneHpcCacheTaskInfoRespStatus
	FAIL    OneHpcCacheTaskInfoRespStatus
}

func GetOneHpcCacheTaskInfoRespStatusEnum() OneHpcCacheTaskInfoRespStatusEnum {
	return OneHpcCacheTaskInfoRespStatusEnum{
		SUCCESS: OneHpcCacheTaskInfoRespStatus{
			value: "SUCCESS",
		},
		DOING: OneHpcCacheTaskInfoRespStatus{
			value: "DOING",
		},
		FAIL: OneHpcCacheTaskInfoRespStatus{
			value: "FAIL",
		},
	}
}

func (c OneHpcCacheTaskInfoRespStatus) Value() string {
	return c.value
}

func (c OneHpcCacheTaskInfoRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneHpcCacheTaskInfoRespStatus) UnmarshalJSON(b []byte) error {
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
