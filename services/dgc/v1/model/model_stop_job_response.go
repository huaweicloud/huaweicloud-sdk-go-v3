package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobResponse struct{}"
	}

	return strings.Join([]string{"StopJobResponse", string(data)}, " ")
}
