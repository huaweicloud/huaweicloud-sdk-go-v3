package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEngressEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEngressEipV2Response struct{}"
	}

	return strings.Join([]string{"UpdateEngressEipV2Response", string(data)}, " ")
}
