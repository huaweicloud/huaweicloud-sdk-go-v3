package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopIndividualStreamJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopIndividualStreamJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopIndividualStreamJobResponse struct{}"
	}

	return strings.Join([]string{"StopIndividualStreamJobResponse", string(data)}, " ")
}
