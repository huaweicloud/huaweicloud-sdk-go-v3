package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIpsWhitelistResponse Response Object
type BatchCreateIpsWhitelistResponse struct {
	Data           *IpsWhitelistRespData `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchCreateIpsWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpsWhitelistResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateIpsWhitelistResponse", string(data)}, " ")
}
