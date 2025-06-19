package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableIpBlacklistResponse Response Object
type EnableIpBlacklistResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableIpBlacklistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableIpBlacklistResponse struct{}"
	}

	return strings.Join([]string{"EnableIpBlacklistResponse", string(data)}, " ")
}
