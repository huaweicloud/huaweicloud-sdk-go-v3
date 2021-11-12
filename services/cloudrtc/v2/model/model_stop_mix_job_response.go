package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopMixJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopMixJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMixJobResponse struct{}"
	}

	return strings.Join([]string{"StopMixJobResponse", string(data)}, " ")
}
