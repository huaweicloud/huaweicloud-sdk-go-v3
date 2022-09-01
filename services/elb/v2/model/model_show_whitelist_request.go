package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWhitelistRequest struct {

	// 白名单的id
	WhitelistId string `json:"whitelist_id" xml:"whitelist_id"`
}

func (o ShowWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowWhitelistRequest", string(data)}, " ")
}
