package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIpBlacklistResponse Response Object
type ImportIpBlacklistResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ImportIpBlacklistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIpBlacklistResponse struct{}"
	}

	return strings.Join([]string{"ImportIpBlacklistResponse", string(data)}, " ")
}
