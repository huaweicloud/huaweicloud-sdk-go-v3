package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty" xml:"whitelist"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistResponse", string(data)}, " ")
}
