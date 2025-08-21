package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAverageEvaluationResponse Response Object
type ShowAverageEvaluationResponse struct {

	// **参数解释：** 合并请求id。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 评价平均分。
	AverageEvaluationLevel *float64 `json:"average_evaluation_level,omitempty"`

	// **参数解释：** 单人评价详情。
	Evaluations *[]MergeRequestBaseEvaluationDto `json:"evaluations,omitempty"`

	// **参数解释：** 自定义评价维度详情。
	CustomEvaluations *[]MergeRequestCustomAverageEvaluationDto `json:"custom_evaluations,omitempty"`
	HttpStatusCode    int                                       `json:"-"`
}

func (o ShowAverageEvaluationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAverageEvaluationResponse struct{}"
	}

	return strings.Join([]string{"ShowAverageEvaluationResponse", string(data)}, " ")
}
