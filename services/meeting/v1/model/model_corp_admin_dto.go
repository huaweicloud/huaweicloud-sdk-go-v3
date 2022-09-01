package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CorpAdminDto struct {

	// 企业用户账号。 maxLength：64 minLength：1
	Account string `json:"account" xml:"account"`
}

func (o CorpAdminDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpAdminDto struct{}"
	}

	return strings.Join([]string{"CorpAdminDto", string(data)}, " ")
}
