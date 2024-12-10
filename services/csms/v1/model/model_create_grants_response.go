package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGrantsResponse Response Object
type CreateGrantsResponse struct {
	Data           *GrantData `json:"data,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CreateGrantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGrantsResponse struct{}"
	}

	return strings.Join([]string{"CreateGrantsResponse", string(data)}, " ")
}
