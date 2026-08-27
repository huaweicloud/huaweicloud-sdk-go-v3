package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTaskResultCommitRequest Request Object
type ShowTaskResultCommitRequest struct {

	// **参数解释**： 演化任务标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	EvolveTaskId string `json:"evolve_task_id"`

	// **参数解释**： 演化任务结果的commit_id。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-256]个字符。 **默认取值**： 不涉及
	CommitId string `json:"commit_id"`

	// **参数解释**： 从哪个轮次开始查询。 **约束限制**： 不涉及 **取值范围**： [-1-10000]。 **默认取值**： 不涉及
	Iteration int32 `json:"iteration"`

	// **参数解释**： 信息类型。 **约束限制**： 不涉及 **取值范围**： * CODE: 代码文件 * INSIGHT: LLM的见解 * SUMMARY: 结果的汇总指标 **默认取值**： 不涉及
	Type *ShowTaskResultCommitRequestType `json:"type,omitempty"`

	// **参数解释**： 文件路径，默认值为算法文件，算法文件为空时，会自动去INSIGHT返回。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-896]个字符。 **默认取值**： 不涉及
	FilePath *string `json:"file_path,omitempty"`
}

func (o ShowTaskResultCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResultCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskResultCommitRequest", string(data)}, " ")
}

type ShowTaskResultCommitRequestType struct {
	value string
}

type ShowTaskResultCommitRequestTypeEnum struct {
	SUMMARY ShowTaskResultCommitRequestType
	INSIGHT ShowTaskResultCommitRequestType
	CODE    ShowTaskResultCommitRequestType
}

func GetShowTaskResultCommitRequestTypeEnum() ShowTaskResultCommitRequestTypeEnum {
	return ShowTaskResultCommitRequestTypeEnum{
		SUMMARY: ShowTaskResultCommitRequestType{
			value: "SUMMARY",
		},
		INSIGHT: ShowTaskResultCommitRequestType{
			value: "INSIGHT",
		},
		CODE: ShowTaskResultCommitRequestType{
			value: "CODE",
		},
	}
}

func (c ShowTaskResultCommitRequestType) Value() string {
	return c.value
}

func (c ShowTaskResultCommitRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResultCommitRequestType) UnmarshalJSON(b []byte) error {
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
