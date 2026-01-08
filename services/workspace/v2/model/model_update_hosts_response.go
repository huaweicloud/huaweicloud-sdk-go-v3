package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostsResponse Response Object
type UpdateHostsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostsResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostsResponse", string(data)}, " ")
}
