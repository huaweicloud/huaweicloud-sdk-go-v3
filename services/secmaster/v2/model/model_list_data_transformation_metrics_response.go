package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataTransformationMetricsResponse Response Object
type ListDataTransformationMetricsResponse struct {
	Status *Object1Map `json:"status,omitempty"`

	Usage          *Object2Map `json:"usage,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListDataTransformationMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataTransformationMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListDataTransformationMetricsResponse", string(data)}, " ")
}
