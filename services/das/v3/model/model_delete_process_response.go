package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteProcessResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProcessResponse struct{}"
	}

	return strings.Join([]string{"DeleteProcessResponse", string(data)}, " ")
}
