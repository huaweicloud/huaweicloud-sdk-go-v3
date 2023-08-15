package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpResponse Response Object
type CreateIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpResponse struct{}"
	}

	return strings.Join([]string{"CreateIpResponse", string(data)}, " ")
}
