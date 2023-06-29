package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePerformanceResourceResponse Response Object
type UpdatePerformanceResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePerformanceResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePerformanceResourceResponse struct{}"
	}

	return strings.Join([]string{"UpdatePerformanceResourceResponse", string(data)}, " ")
}
