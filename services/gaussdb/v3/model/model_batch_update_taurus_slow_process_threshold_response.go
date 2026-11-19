package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTaurusSlowProcessThresholdResponse Response Object
type BatchUpdateTaurusSlowProcessThresholdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateTaurusSlowProcessThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTaurusSlowProcessThresholdResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateTaurusSlowProcessThresholdResponse", string(data)}, " ")
}
