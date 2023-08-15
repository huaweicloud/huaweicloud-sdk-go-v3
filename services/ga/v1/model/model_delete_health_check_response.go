package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthCheckResponse Response Object
type DeleteHealthCheckResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthCheckResponse", string(data)}, " ")
}
