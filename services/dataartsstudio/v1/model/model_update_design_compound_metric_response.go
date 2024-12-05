package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignCompoundMetricResponse Response Object
type UpdateDesignCompoundMetricResponse struct {
	Data           *UpdateDesignCompoundMetricResultData `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o UpdateDesignCompoundMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignCompoundMetricResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignCompoundMetricResponse", string(data)}, " ")
}
