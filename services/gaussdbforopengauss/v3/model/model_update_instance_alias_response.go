package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceAliasResponse Response Object
type UpdateInstanceAliasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAliasResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAliasResponse", string(data)}, " ")
}
