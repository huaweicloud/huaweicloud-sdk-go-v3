package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePerformanceResourceResponse Response Object
type CreatePerformanceResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePerformanceResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePerformanceResourceResponse struct{}"
	}

	return strings.Join([]string{"CreatePerformanceResourceResponse", string(data)}, " ")
}
