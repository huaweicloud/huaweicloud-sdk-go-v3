package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteApiByVersionIdV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiByVersionIdV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiByVersionIdV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiByVersionIdV2Response", string(data)}, " ")
}
