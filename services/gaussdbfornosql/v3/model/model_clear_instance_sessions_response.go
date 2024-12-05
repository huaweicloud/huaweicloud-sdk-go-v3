package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearInstanceSessionsResponse Response Object
type ClearInstanceSessionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ClearInstanceSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearInstanceSessionsResponse struct{}"
	}

	return strings.Join([]string{"ClearInstanceSessionsResponse", string(data)}, " ")
}
