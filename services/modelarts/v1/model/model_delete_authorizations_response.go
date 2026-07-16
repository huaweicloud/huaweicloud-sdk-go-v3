package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuthorizationsResponse Response Object
type DeleteAuthorizationsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuthorizationsResponse", string(data)}, " ")
}
