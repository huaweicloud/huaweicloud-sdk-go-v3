package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddEngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AddEngressEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEngressEipV2Response struct{}"
	}

	return strings.Join([]string{"AddEngressEipV2Response", string(data)}, " ")
}
