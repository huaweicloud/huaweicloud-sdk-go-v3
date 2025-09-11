package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WhitelistBatchDeleteRequest struct {

	// 白名单记录ID
	Whitelists []string `json:"whitelists"`
}

func (o WhitelistBatchDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhitelistBatchDeleteRequest struct{}"
	}

	return strings.Join([]string{"WhitelistBatchDeleteRequest", string(data)}, " ")
}
