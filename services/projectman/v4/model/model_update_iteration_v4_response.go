package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateIterationV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateIterationV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIterationV4Response struct{}"
	}

	return strings.Join([]string{"UpdateIterationV4Response", string(data)}, " ")
}
