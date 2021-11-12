package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAppV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppV2Response struct{}"
	}

	return strings.Join([]string{"DeleteAppV2Response", string(data)}, " ")
}
