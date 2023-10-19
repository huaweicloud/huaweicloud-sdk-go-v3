package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionNamesResponse Response Object
type ListFunctionNamesResponse struct {
	Body *[]string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFunctionNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionNamesResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionNamesResponse", string(data)}, " ")
}
