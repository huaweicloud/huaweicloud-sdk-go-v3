package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTuningResponse Response Object
type ShowTuningResponse struct {
	TuneResult     *AdviceResult `json:"tune_result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowTuningResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTuningResponse struct{}"
	}

	return strings.Join([]string{"ShowTuningResponse", string(data)}, " ")
}
