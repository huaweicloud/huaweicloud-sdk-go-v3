package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CorpAdminDto struct {
	// 企业用户账号。 maxLength：64 minLength：1

	Account string `json:"account"`
}

func (o CorpAdminDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpAdminDto struct{}"
	}

	return strings.Join([]string{"CorpAdminDto", string(data)}, " ")
}
