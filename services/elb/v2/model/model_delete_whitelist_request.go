package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWhitelistRequest Request Object
type DeleteWhitelistRequest struct {

	// 白名单id
	WhitelistId string `json:"whitelist_id"`
}

func (o DeleteWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWhitelistRequest struct{}"
	}

	return strings.Join([]string{"DeleteWhitelistRequest", string(data)}, " ")
}
