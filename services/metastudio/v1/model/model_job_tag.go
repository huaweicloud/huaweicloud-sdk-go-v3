package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobTag 任务标签。 * ECOMMERCE: 电商 * NEWS: 新闻 * MARKETING: 营销 * LIVE: 直播 * EDUCATION: 教培 * CUSTOMER: 客服 * STORYTELLING: 故事
type JobTag struct {
	value string
}

type JobTagEnum struct {
	ECOMMERCE    JobTag
	NEWS         JobTag
	MARKETING    JobTag
	LIVE         JobTag
	EDUCATION    JobTag
	CUSTOMER     JobTag
	STORYTELLING JobTag
}

func GetJobTagEnum() JobTagEnum {
	return JobTagEnum{
		ECOMMERCE: JobTag{
			value: "ECOMMERCE",
		},
		NEWS: JobTag{
			value: "NEWS",
		},
		MARKETING: JobTag{
			value: "MARKETING",
		},
		LIVE: JobTag{
			value: "LIVE",
		},
		EDUCATION: JobTag{
			value: "EDUCATION",
		},
		CUSTOMER: JobTag{
			value: "CUSTOMER",
		},
		STORYTELLING: JobTag{
			value: "STORYTELLING",
		},
	}
}

func (c JobTag) Value() string {
	return c.value
}

func (c JobTag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobTag) UnmarshalJSON(b []byte) error {
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
