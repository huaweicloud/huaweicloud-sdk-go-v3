package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWhitelistResponse Response Object
type UpdateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistResponse", string(data)}, " ")
}
