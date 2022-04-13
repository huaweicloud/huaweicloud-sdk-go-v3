package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAppVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppVersionResponse", string(data)}, " ")
}
