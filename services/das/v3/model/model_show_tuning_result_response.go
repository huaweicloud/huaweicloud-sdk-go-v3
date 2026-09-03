package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTuningResultResponse Response Object
type ShowTuningResultResponse struct {
	TuneResult     *AdviceResult `json:"tune_result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowTuningResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTuningResultResponse struct{}"
	}

	return strings.Join([]string{"ShowTuningResultResponse", string(data)}, " ")
}
