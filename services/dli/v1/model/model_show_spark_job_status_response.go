package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSparkJobStatusResponse Response Object
type ShowSparkJobStatusResponse struct {

	// 参数解释:  批处理作业的ID，采用UUID（通用唯一识别码）格式 示例: 0a324461-d9d9-45da-a52a-3b3c7a3d809e 约束限制:  无 取值范围: 无 默认取值: 无
	Id *string `json:"id,omitempty"`

	// 参数解释:  批处理作业的状态。starting：正在启动；running：正在执行任务；dead：session已退出；success：session停止成功；recovering：正在恢复 示例: success 约束限制:  无 取值范围: starting（批处理作业正在启动） running（批处理作业正在执行任务） dead（批处理作业已退出） success（批处理作业执行成功） recovering（批处理作业正在恢复） 默认取值: 无
	State          *ShowSparkJobStatusResponseState `json:"state,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowSparkJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobStatusResponse", string(data)}, " ")
}

type ShowSparkJobStatusResponseState struct {
	value string
}

type ShowSparkJobStatusResponseStateEnum struct {
	STARTING   ShowSparkJobStatusResponseState
	RUNNING    ShowSparkJobStatusResponseState
	DEAD       ShowSparkJobStatusResponseState
	SUCCESS    ShowSparkJobStatusResponseState
	RECOVERING ShowSparkJobStatusResponseState
}

func GetShowSparkJobStatusResponseStateEnum() ShowSparkJobStatusResponseStateEnum {
	return ShowSparkJobStatusResponseStateEnum{
		STARTING: ShowSparkJobStatusResponseState{
			value: "starting",
		},
		RUNNING: ShowSparkJobStatusResponseState{
			value: "running",
		},
		DEAD: ShowSparkJobStatusResponseState{
			value: "dead",
		},
		SUCCESS: ShowSparkJobStatusResponseState{
			value: "success",
		},
		RECOVERING: ShowSparkJobStatusResponseState{
			value: "recovering",
		},
	}
}

func (c ShowSparkJobStatusResponseState) Value() string {
	return c.value
}

func (c ShowSparkJobStatusResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSparkJobStatusResponseState) UnmarshalJSON(b []byte) error {
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
