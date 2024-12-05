package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignCompoundMetricResponse Response Object
type DeleteDesignCompoundMetricResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignCompoundMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignCompoundMetricResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignCompoundMetricResponse", string(data)}, " ")
}
