package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetricDataResponse Response Object
type CreateMetricDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetricDataResponse struct{}"
	}

	return strings.Join([]string{"CreateMetricDataResponse", string(data)}, " ")
}
