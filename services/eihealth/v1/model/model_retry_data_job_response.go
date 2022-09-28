package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RetryDataJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryDataJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryDataJobResponse struct{}"
	}

	return strings.Join([]string{"RetryDataJobResponse", string(data)}, " ")
}
