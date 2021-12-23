package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSharedTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSharedTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSharedTagResponse", string(data)}, " ")
}
