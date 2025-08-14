package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAliasResponse Response Object
type CreateAliasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAliasResponse struct{}"
	}

	return strings.Join([]string{"CreateAliasResponse", string(data)}, " ")
}
