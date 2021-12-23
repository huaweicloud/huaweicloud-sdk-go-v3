package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteStreamV3Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStreamV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamV3Response struct{}"
	}

	return strings.Join([]string{"DeleteStreamV3Response", string(data)}, " ")
}
