package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstancesPeriodOrderNewRequest Request Object
type CreateInstancesPeriodOrderNewRequest struct {
	Body *CreateInstancePeriodRequest `json:"body,omitempty"`
}

func (o CreateInstancesPeriodOrderNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancesPeriodOrderNewRequest struct{}"
	}

	return strings.Join([]string{"CreateInstancesPeriodOrderNewRequest", string(data)}, " ")
}
