package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateWhitelistRequest struct {

	// 待更新的白名单id
	WhitelistId string `json:"whitelist_id"`

	Body *UpdateWhitelistRequestBody `json:"body,omitempty"`
}

func (o UpdateWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequest", string(data)}, " ")
}
