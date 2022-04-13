package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
