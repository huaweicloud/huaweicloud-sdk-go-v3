package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSearchMetricHitsResponse Response Object
type BatchSearchMetricHitsResponse struct {
	Body           *[]ShowMetricResultResponseBody `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o BatchSearchMetricHitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSearchMetricHitsResponse struct{}"
	}

	return strings.Join([]string{"BatchSearchMetricHitsResponse", string(data)}, " ")
}
