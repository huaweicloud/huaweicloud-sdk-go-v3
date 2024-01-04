package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryTaskCompletionResponse Response Object
type ListFactoryTaskCompletionResponse struct {

	// 昨天的任务信息
	Yesterday *[]ListFactoryTaskCompletionResYesterday `json:"yesterday,omitempty"`

	// 近7天的平均任务信息
	Average *[]ListFactoryTaskCompletionResAverage `json:"average,omitempty"`

	// 当天的任务信息
	Today          *[]ListFactoryTaskCompletionResToday `json:"today,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListFactoryTaskCompletionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryTaskCompletionResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryTaskCompletionResponse", string(data)}, " ")
}
