package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpsWhitelistResponse Response Object
type UpdateIpsWhitelistResponse struct {
	Data           *UpdateIpsWhitelistRespData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateIpsWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpsWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpsWhitelistResponse", string(data)}, " ")
}
