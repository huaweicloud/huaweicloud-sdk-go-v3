package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIpsWhitelistResponse Response Object
type BatchDeleteIpsWhitelistResponse struct {
	Data           *IpsWhitelistRespData `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchDeleteIpsWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpsWhitelistResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpsWhitelistResponse", string(data)}, " ")
}
