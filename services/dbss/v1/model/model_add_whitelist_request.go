package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWhitelistRequest 添加白名单请求体
type AddWhitelistRequest struct {

	// SQL白名单列表
	Whitelists []ObjWhitelist `json:"whitelists"`
}

func (o AddWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWhitelistRequest struct{}"
	}

	return strings.Join([]string{"AddWhitelistRequest", string(data)}, " ")
}
