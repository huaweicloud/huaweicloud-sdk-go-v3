package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppV3Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAppV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppV3Response struct{}"
	}

	return strings.Join([]string{"CreateAppV3Response", string(data)}, " ")
}
