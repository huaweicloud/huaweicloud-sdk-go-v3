package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeApiToInstanceResponse Response Object
type AuthorizeApiToInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AuthorizeApiToInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeApiToInstanceResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeApiToInstanceResponse", string(data)}, " ")
}
