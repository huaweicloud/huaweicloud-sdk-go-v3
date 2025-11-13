package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FeedbackType **参数解释**： 反馈类型。 **约束限制**： 不涉及 **取值范围**： * upvote：点赞 * downvote：点踩 * none：取消反馈 **默认取值**： 不涉及
type FeedbackType struct {
	value string
}

type FeedbackTypeEnum struct {
	UPVOTE   FeedbackType
	DOWNVOTE FeedbackType
	NONE     FeedbackType
}

func GetFeedbackTypeEnum() FeedbackTypeEnum {
	return FeedbackTypeEnum{
		UPVOTE: FeedbackType{
			value: "upvote",
		},
		DOWNVOTE: FeedbackType{
			value: "downvote",
		},
		NONE: FeedbackType{
			value: "none",
		},
	}
}

func (c FeedbackType) Value() string {
	return c.value
}

func (c FeedbackType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FeedbackType) UnmarshalJSON(b []byte) error {
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
