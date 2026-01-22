package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryIpBlacklistResponse Response Object
type RetryIpBlacklistResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RetryIpBlacklistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryIpBlacklistResponse struct{}"
	}

	return strings.Join([]string{"RetryIpBlacklistResponse", string(data)}, " ")
}
