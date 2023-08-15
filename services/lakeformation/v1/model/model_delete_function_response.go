package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionResponse Response Object
type DeleteFunctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionResponse", string(data)}, " ")
}
