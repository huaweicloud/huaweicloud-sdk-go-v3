package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppIconResponse Response Object
type DeleteAppIconResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppIconResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppIconResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppIconResponse", string(data)}, " ")
}
