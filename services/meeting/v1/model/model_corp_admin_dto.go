package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CorpAdminDto struct {

	// 企业用户帐号。 * 如果是帐号/密码鉴权方式，是指华为云会议帐号 * 如果是App ID鉴权方式，是指第三方User ID
	Account string `json:"account"`
}

func (o CorpAdminDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpAdminDto struct{}"
	}

	return strings.Join([]string{"CorpAdminDto", string(data)}, " ")
}
