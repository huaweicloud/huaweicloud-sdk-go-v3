package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobActionType 对训练作业的操作请求体。终止训练作业请使用terminate。
type JobActionType struct {

	// 对训练作业的操作请求。参数值设置为terminate时，表示终止训练作业操作。
	ActionType string `json:"action_type"`
}

func (o JobActionType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobActionType struct{}"
	}

	return strings.Join([]string{"JobActionType", string(data)}, " ")
}
