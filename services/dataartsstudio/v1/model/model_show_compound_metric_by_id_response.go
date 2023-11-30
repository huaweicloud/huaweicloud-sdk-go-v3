package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompoundMetricByIdResponse Response Object
type ShowCompoundMetricByIdResponse struct {
	Data           *CompoundMetricVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowCompoundMetricByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompoundMetricByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowCompoundMetricByIdResponse", string(data)}, " ")
}
