package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopJobInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopJobInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopJobInstanceResponse", string(data)}, " ")
}
