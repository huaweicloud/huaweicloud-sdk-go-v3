package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteVersionAliasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVersionAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVersionAliasResponse struct{}"
	}

	return strings.Join([]string{"DeleteVersionAliasResponse", string(data)}, " ")
}
