package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PutMergeRequestParamsDto struct {

	// 合并请求标题
	Title *string `json:"title,omitempty"`

	// 合并请求状态
	StateEvent *PutMergeRequestParamsDtoStateEvent `json:"state_event,omitempty"`

	// 合并人id列表
	AssigneeIds *string `json:"assignee_ids,omitempty"`

	// 评审人id列表
	ReviewerIds *string `json:"reviewer_ids,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 里程碑id
	MilestoneId *int32 `json:"milestone_id,omitempty"`

	// 标签
	Labels *interface{} `json:"labels,omitempty"`

	// 合入后删除源分支
	ForceRemoveSourceBranch *bool `json:"force_remove_source_branch,omitempty"`

	// squash合入
	Squash *bool `json:"squash,omitempty"`

	// squash提交信息
	SquashCommitMessage *string `json:"squash_commit_message,omitempty"`

	// e2e单号
	WorkItemIds *[]string `json:"work_item_ids,omitempty"`
}

func (o PutMergeRequestParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutMergeRequestParamsDto struct{}"
	}

	return strings.Join([]string{"PutMergeRequestParamsDto", string(data)}, " ")
}

type PutMergeRequestParamsDtoStateEvent struct {
	value string
}

type PutMergeRequestParamsDtoStateEventEnum struct {
	REOPEN PutMergeRequestParamsDtoStateEvent
	CLOSE  PutMergeRequestParamsDtoStateEvent
}

func GetPutMergeRequestParamsDtoStateEventEnum() PutMergeRequestParamsDtoStateEventEnum {
	return PutMergeRequestParamsDtoStateEventEnum{
		REOPEN: PutMergeRequestParamsDtoStateEvent{
			value: "reopen",
		},
		CLOSE: PutMergeRequestParamsDtoStateEvent{
			value: "close",
		},
	}
}

func (c PutMergeRequestParamsDtoStateEvent) Value() string {
	return c.value
}

func (c PutMergeRequestParamsDtoStateEvent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutMergeRequestParamsDtoStateEvent) UnmarshalJSON(b []byte) error {
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
