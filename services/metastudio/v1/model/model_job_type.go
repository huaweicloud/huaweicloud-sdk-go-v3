package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobType 训练类型。 * BASIC: 基础版(20句话) * MIDDLE: 进阶版(100句话) * ADVANCE: 高级版 * THIRD_PARTY: 第三方出门问问训练版 * THIRD_PARTY_LJZN: 第三方逻辑智能训练版 * FLEXUS: Flexus版---用的是大模型特征提取
type JobType struct {
	value string
}

type JobTypeEnum struct {
	BASIC            JobType
	MIDDLE           JobType
	ADVANCE          JobType
	THIRD_PARTY      JobType
	THIRD_PARTY_LJZN JobType
	FLEXUS           JobType
}

func GetJobTypeEnum() JobTypeEnum {
	return JobTypeEnum{
		BASIC: JobType{
			value: "BASIC",
		},
		MIDDLE: JobType{
			value: "MIDDLE",
		},
		ADVANCE: JobType{
			value: "ADVANCE",
		},
		THIRD_PARTY: JobType{
			value: "THIRD_PARTY",
		},
		THIRD_PARTY_LJZN: JobType{
			value: "THIRD_PARTY_LJZN",
		},
		FLEXUS: JobType{
			value: "FLEXUS",
		},
	}
}

func (c JobType) Value() string {
	return c.value
}

func (c JobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobType) UnmarshalJSON(b []byte) error {
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
