package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppInstanceResponse Response Object
type DeleteAppInstanceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAppInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppInstanceResponse", string(data)}, " ")
}
