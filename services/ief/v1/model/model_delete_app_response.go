package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppResponse Response Object
type DeleteAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppResponse", string(data)}, " ")
}
