package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeDataConnectionResponse Response Object
type AuthorizeDataConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AuthorizeDataConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeDataConnectionResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeDataConnectionResponse", string(data)}, " ")
}
