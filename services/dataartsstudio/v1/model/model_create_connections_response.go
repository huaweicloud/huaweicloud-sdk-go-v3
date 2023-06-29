package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionsResponse Response Object
type CreateConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionsResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionsResponse", string(data)}, " ")
}
