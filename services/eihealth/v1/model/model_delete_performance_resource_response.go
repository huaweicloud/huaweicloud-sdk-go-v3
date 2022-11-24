package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePerformanceResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePerformanceResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePerformanceResourceResponse struct{}"
	}

	return strings.Join([]string{"DeletePerformanceResourceResponse", string(data)}, " ")
}
