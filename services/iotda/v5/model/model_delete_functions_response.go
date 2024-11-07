package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionsResponse Response Object
type DeleteFunctionsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFunctionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionsResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionsResponse", string(data)}, " ")
}
