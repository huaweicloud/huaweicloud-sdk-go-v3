package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTaskRunningDetailsRequest Request Object
type ShowTaskRunningDetailsRequest struct {

	// **参数解释**： 演化任务标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	EvolveTaskId string `json:"evolve_task_id"`

	// **参数解释**： 统计信息类型。 **约束限制**： 不涉及 **取值范围**： * PROGRESS： 进度信息。 * SUMMARY:   结果统计值。 * BEST_RESULT:  最优结果的commitId。 * GENERATION_STATS: 各迭代的统计值。 **默认取值**： 不涉及
	Type ShowTaskRunningDetailsRequestType `json:"type"`
}

func (o ShowTaskRunningDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRunningDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRunningDetailsRequest", string(data)}, " ")
}

type ShowTaskRunningDetailsRequestType struct {
	value string
}

type ShowTaskRunningDetailsRequestTypeEnum struct {
	PROGRESS         ShowTaskRunningDetailsRequestType
	SUMMARY          ShowTaskRunningDetailsRequestType
	BEST_RESULT      ShowTaskRunningDetailsRequestType
	GENERATION_STATS ShowTaskRunningDetailsRequestType
}

func GetShowTaskRunningDetailsRequestTypeEnum() ShowTaskRunningDetailsRequestTypeEnum {
	return ShowTaskRunningDetailsRequestTypeEnum{
		PROGRESS: ShowTaskRunningDetailsRequestType{
			value: "progress",
		},
		SUMMARY: ShowTaskRunningDetailsRequestType{
			value: "summary",
		},
		BEST_RESULT: ShowTaskRunningDetailsRequestType{
			value: "best_result",
		},
		GENERATION_STATS: ShowTaskRunningDetailsRequestType{
			value: "generation_stats",
		},
	}
}

func (c ShowTaskRunningDetailsRequestType) Value() string {
	return c.value
}

func (c ShowTaskRunningDetailsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskRunningDetailsRequestType) UnmarshalJSON(b []byte) error {
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
