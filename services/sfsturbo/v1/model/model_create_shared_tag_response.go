package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSharedTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSharedTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSharedTagResponse", string(data)}, " ")
}
