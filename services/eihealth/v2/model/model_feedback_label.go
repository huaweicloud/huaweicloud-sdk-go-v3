package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FeedbackLabel **参数解释**： 反馈标签。 **约束限制**： 不涉及 **取值范围**： * others：其他 * professional：内容专业 * helpful：有帮助 * unprofessional：不专业 * unhelpful ：没有帮助 * unsafe：有害/不安全 * misinformation：虚假信息 * download_successful：下载数据成功 * promising：生成分子有价值 * synthesizable：合成可行性高 * validated：运行结果与实验相符 * download_failed：下载数据失败 * unpromising：生成分子没价值 * hard_to_synthesize：较难合成 * not_validated：运行结果与实验相悖 * plan_validated：规划步骤与预期相符 * accurate_results：执行结果准确 * high_performance：运行速度快 * plan_flawed：规划步骤不合理 * inaccurate_results：选择算法不准确 * low_performance：运行速度慢 **默认取值**： 不涉及
type FeedbackLabel struct {
	value string
}

type FeedbackLabelEnum struct {
	OTHERS              FeedbackLabel
	PROFESSIONAL        FeedbackLabel
	HELPFUL             FeedbackLabel
	UNPROFESSIONAL      FeedbackLabel
	UNHELPFUL           FeedbackLabel
	UNSAFE              FeedbackLabel
	MISINFORMATION      FeedbackLabel
	DOWNLOAD_SUCCESSFUL FeedbackLabel
	PROMISING           FeedbackLabel
	SYNTHESIZABLE       FeedbackLabel
	VALIDATED           FeedbackLabel
	DOWNLOAD_FAILED     FeedbackLabel
	UNPROMISING         FeedbackLabel
	HARD_TO_SYNTHESIZE  FeedbackLabel
	NOT_VALIDATED       FeedbackLabel
	PLAN_VALIDATED      FeedbackLabel
	ACCURATE_RESULTS    FeedbackLabel
	HIGH_PERFORMANCE    FeedbackLabel
	PLAN_FLAWED         FeedbackLabel
	INACCURATE_RESULTS  FeedbackLabel
	LOW_PERFORMANCE     FeedbackLabel
}

func GetFeedbackLabelEnum() FeedbackLabelEnum {
	return FeedbackLabelEnum{
		OTHERS: FeedbackLabel{
			value: "others",
		},
		PROFESSIONAL: FeedbackLabel{
			value: "professional",
		},
		HELPFUL: FeedbackLabel{
			value: "helpful",
		},
		UNPROFESSIONAL: FeedbackLabel{
			value: "unprofessional",
		},
		UNHELPFUL: FeedbackLabel{
			value: "unhelpful",
		},
		UNSAFE: FeedbackLabel{
			value: "unsafe",
		},
		MISINFORMATION: FeedbackLabel{
			value: "misinformation",
		},
		DOWNLOAD_SUCCESSFUL: FeedbackLabel{
			value: "download_successful",
		},
		PROMISING: FeedbackLabel{
			value: "promising",
		},
		SYNTHESIZABLE: FeedbackLabel{
			value: "synthesizable",
		},
		VALIDATED: FeedbackLabel{
			value: "validated",
		},
		DOWNLOAD_FAILED: FeedbackLabel{
			value: "download_failed",
		},
		UNPROMISING: FeedbackLabel{
			value: "unpromising",
		},
		HARD_TO_SYNTHESIZE: FeedbackLabel{
			value: "hard_to_synthesize",
		},
		NOT_VALIDATED: FeedbackLabel{
			value: "not_validated",
		},
		PLAN_VALIDATED: FeedbackLabel{
			value: "plan_validated",
		},
		ACCURATE_RESULTS: FeedbackLabel{
			value: "accurate_results",
		},
		HIGH_PERFORMANCE: FeedbackLabel{
			value: "high_performance",
		},
		PLAN_FLAWED: FeedbackLabel{
			value: "plan_flawed",
		},
		INACCURATE_RESULTS: FeedbackLabel{
			value: "inaccurate_results",
		},
		LOW_PERFORMANCE: FeedbackLabel{
			value: "low_performance",
		},
	}
}

func (c FeedbackLabel) Value() string {
	return c.value
}

func (c FeedbackLabel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FeedbackLabel) UnmarshalJSON(b []byte) error {
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
