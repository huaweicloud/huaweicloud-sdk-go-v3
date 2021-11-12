package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"CreateWhitelistResponse", string(data)}, " ")
}
