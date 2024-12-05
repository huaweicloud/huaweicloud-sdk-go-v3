package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignCompoundMetricResponse Response Object
type CreateDesignCompoundMetricResponse struct {
	Data           *CreateDesignCompoundMetricResultData `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o CreateDesignCompoundMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignCompoundMetricResponse struct{}"
	}

	return strings.Join([]string{"CreateDesignCompoundMetricResponse", string(data)}, " ")
}
