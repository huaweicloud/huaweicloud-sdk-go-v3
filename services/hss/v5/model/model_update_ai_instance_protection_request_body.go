package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAiInstanceProtectionRequestBody 修改OpenClaw防护状态请求结构体
type UpdateAiInstanceProtectionRequestBody struct {

	// 需要修改防护状态的OpenClaw智能体ID列表 **约束限制**： - 单次最多50个 **取值范围**： 智能体实例的agent_id
	AgentIdList []string `json:"agent_id_list"`

	// **参数解释**: 目标防护状态 **取值范围**: - open：开启防护 - close：关闭防
	Status UpdateAiInstanceProtectionRequestBodyStatus `json:"status"`
}

func (o UpdateAiInstanceProtectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiInstanceProtectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAiInstanceProtectionRequestBody", string(data)}, " ")
}

type UpdateAiInstanceProtectionRequestBodyStatus struct {
	value string
}

type UpdateAiInstanceProtectionRequestBodyStatusEnum struct {
	OPEN  UpdateAiInstanceProtectionRequestBodyStatus
	CLOSE UpdateAiInstanceProtectionRequestBodyStatus
}

func GetUpdateAiInstanceProtectionRequestBodyStatusEnum() UpdateAiInstanceProtectionRequestBodyStatusEnum {
	return UpdateAiInstanceProtectionRequestBodyStatusEnum{
		OPEN: UpdateAiInstanceProtectionRequestBodyStatus{
			value: "open",
		},
		CLOSE: UpdateAiInstanceProtectionRequestBodyStatus{
			value: "close",
		},
	}
}

func (c UpdateAiInstanceProtectionRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateAiInstanceProtectionRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiInstanceProtectionRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
