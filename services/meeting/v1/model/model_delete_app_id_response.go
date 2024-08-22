package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppIdResponse Response Object
type DeleteAppIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppIdResponse", string(data)}, " ")
}
