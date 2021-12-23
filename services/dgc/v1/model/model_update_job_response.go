package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}
