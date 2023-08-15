package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSynthesisTaskResultResponse Response Object
type ShowSynthesisTaskResultResponse struct {
	Status *AsyncTaskStatus `json:"status,omitempty"`

	TaskData *SynthesisTaskData `json:"task_data,omitempty"`

	Result         *SynthesisResult `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSynthesisTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSynthesisTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSynthesisTaskResultResponse", string(data)}, " ")
}
