package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicKibanaWhitelistReq struct {
	// 白名单。

	WhiteList string `json:"whiteList"`
}

func (o UpdatePublicKibanaWhitelistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicKibanaWhitelistReq struct{}"
	}

	return strings.Join([]string{"UpdatePublicKibanaWhitelistReq", string(data)}, " ")
}
