package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseQueryAssessTaskListResponseData struct {

	// 总任务数
	Total *int32 `json:"total,omitempty"`

	// 评估任务数据列表
	Result *[]QueryAssessTaskResponse `json:"result,omitempty"`
}

func (o BaseQueryAssessTaskListResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseQueryAssessTaskListResponseData struct{}"
	}

	return strings.Join([]string{"BaseQueryAssessTaskListResponseData", string(data)}, " ")
}
