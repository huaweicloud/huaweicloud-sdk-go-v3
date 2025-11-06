package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MergeRequestVoteBodyDto 更新合并请求打分请求体
type MergeRequestVoteBodyDto struct {

	// **参数解释：** 分数。
	Score *int32 `json:"score,omitempty"`

	// **参数解释：** 操作类型(同一分数再次调用会删除打分, 传vote则不会删除)。
	Action *MergeRequestVoteBodyDtoAction `json:"action,omitempty"`
}

func (o MergeRequestVoteBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestVoteBodyDto struct{}"
	}

	return strings.Join([]string{"MergeRequestVoteBodyDto", string(data)}, " ")
}

type MergeRequestVoteBodyDtoAction struct {
	value string
}

type MergeRequestVoteBodyDtoActionEnum struct {
	VOTE MergeRequestVoteBodyDtoAction
}

func GetMergeRequestVoteBodyDtoActionEnum() MergeRequestVoteBodyDtoActionEnum {
	return MergeRequestVoteBodyDtoActionEnum{
		VOTE: MergeRequestVoteBodyDtoAction{
			value: "vote",
		},
	}
}

func (c MergeRequestVoteBodyDtoAction) Value() string {
	return c.value
}

func (c MergeRequestVoteBodyDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MergeRequestVoteBodyDtoAction) UnmarshalJSON(b []byte) error {
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
