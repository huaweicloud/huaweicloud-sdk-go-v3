package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShareConnectionsResponse Response Object
type CreateShareConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateShareConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareConnectionsResponse struct{}"
	}

	return strings.Join([]string{"CreateShareConnectionsResponse", string(data)}, " ")
}
