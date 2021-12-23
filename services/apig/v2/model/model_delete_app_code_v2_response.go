package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAppCodeV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppCodeV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppCodeV2Response struct{}"
	}

	return strings.Join([]string{"DeleteAppCodeV2Response", string(data)}, " ")
}
