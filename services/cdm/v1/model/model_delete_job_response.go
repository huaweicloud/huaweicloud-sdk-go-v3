package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobResponse", string(data)}, " ")
}
