package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateGaussMySqlInstanceAliasResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceAliasResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceAliasResponse", string(data)}, " ")
}
