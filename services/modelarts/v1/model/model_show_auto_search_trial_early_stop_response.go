package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchTrialEarlyStopResponse Response Object
type ShowAutoSearchTrialEarlyStopResponse struct {

	// 提前终止的trial的trial_id。
	EarlystopTrial *string `json:"earlystop_trial,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutoSearchTrialEarlyStopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchTrialEarlyStopResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchTrialEarlyStopResponse", string(data)}, " ")
}
