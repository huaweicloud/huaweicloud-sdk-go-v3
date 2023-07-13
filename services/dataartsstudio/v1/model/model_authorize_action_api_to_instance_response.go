package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeActionApiToInstanceResponse Response Object
type AuthorizeActionApiToInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AuthorizeActionApiToInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeActionApiToInstanceResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeActionApiToInstanceResponse", string(data)}, " ")
}
