package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAliasReqBody struct {

	// 用户自定义的身份源ID对应的别名
	Alias string `json:"alias"`
}

func (o CreateAliasReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAliasReqBody struct{}"
	}

	return strings.Join([]string{"CreateAliasReqBody", string(data)}, " ")
}
