package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompoundMetricVoSearchResultData struct {
	Value *CompoundMetricVoSearchResultDataValue `json:"value,omitempty"`
}

func (o CompoundMetricVoSearchResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompoundMetricVoSearchResultData struct{}"
	}

	return strings.Join([]string{"CompoundMetricVoSearchResultData", string(data)}, " ")
}
