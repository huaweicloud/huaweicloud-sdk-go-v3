package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StageName **参数解释**： 绑定状态。 **约束限制**： 不涉及 **取值范围**： * requirement_analyzer：构建需求文档。 * modeling：构建数学模型。 * data：校验模型数据。 * solver：求解数学模型。 * report：业务辅助分析。 * business_planner：构建需求文档 * data_agent：原始数据处理 * vrp：路径规划求解 * predict_step1：解析任务意图 * predict_step2：探查数据面貌 * predict_step3：生成算法模型 * predict_step4：总结分析报告 **默认取值**： 不涉及
type StageName struct {
	value string
}

type StageNameEnum struct {
	REQUIREMENT_ANALYZER StageName
	MODELING             StageName
	DATA                 StageName
	SOLVER               StageName
	REPORT               StageName
	BUSINESS_PLANNER     StageName
	DATA_AGENT           StageName
	VRP                  StageName
	PREDICT_STEP1        StageName
	PREDICT_STEP2        StageName
	PREDICT_STEP3        StageName
	PREDICT_STEP4        StageName
}

func GetStageNameEnum() StageNameEnum {
	return StageNameEnum{
		REQUIREMENT_ANALYZER: StageName{
			value: "requirement_analyzer",
		},
		MODELING: StageName{
			value: "modeling",
		},
		DATA: StageName{
			value: "data",
		},
		SOLVER: StageName{
			value: "solver",
		},
		REPORT: StageName{
			value: "report",
		},
		BUSINESS_PLANNER: StageName{
			value: "business_planner",
		},
		DATA_AGENT: StageName{
			value: "data_agent",
		},
		VRP: StageName{
			value: "vrp",
		},
		PREDICT_STEP1: StageName{
			value: "predict_step1",
		},
		PREDICT_STEP2: StageName{
			value: "predict_step2",
		},
		PREDICT_STEP3: StageName{
			value: "predict_step3",
		},
		PREDICT_STEP4: StageName{
			value: "predict_step4",
		},
	}
}

func (c StageName) Value() string {
	return c.value
}

func (c StageName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StageName) UnmarshalJSON(b []byte) error {
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
