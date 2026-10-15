package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartJobResp 启动任务返回体。
type StartJobResp struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  启动任务的操作状态，表示批量启动任务中单个任务的执行结果。  **约束限制**：  不涉及。  **取值范围**：  - success：启动成功 - failed：启动失败  **默认取值**：  不涉及
	Status *StartJobRespStatus `json:"status,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o StartJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobResp struct{}"
	}

	return strings.Join([]string{"StartJobResp", string(data)}, " ")
}

type StartJobRespStatus struct {
	value string
}

type StartJobRespStatusEnum struct {
	SUCCESS StartJobRespStatus
	FAILED  StartJobRespStatus
}

func GetStartJobRespStatusEnum() StartJobRespStatusEnum {
	return StartJobRespStatusEnum{
		SUCCESS: StartJobRespStatus{
			value: "success",
		},
		FAILED: StartJobRespStatus{
			value: "failed",
		},
	}
}

func (c StartJobRespStatus) Value() string {
	return c.value
}

func (c StartJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartJobRespStatus) UnmarshalJSON(b []byte) error {
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
