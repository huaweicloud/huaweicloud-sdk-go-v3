package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVhostResponse Response Object
type CreateVhostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVhostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVhostResponse struct{}"
	}

	return strings.Join([]string{"CreateVhostResponse", string(data)}, " ")
}
