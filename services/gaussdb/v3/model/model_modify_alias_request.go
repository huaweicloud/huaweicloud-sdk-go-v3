package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyAliasRequest struct {

	// 实例备注。  长度可在0~64个字符之间，由字母、数字、汉字、英文句号、下划线、中划线组成。
	Alias string `json:"alias"`
}

func (o ModifyAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAliasRequest struct{}"
	}

	return strings.Join([]string{"ModifyAliasRequest", string(data)}, " ")
}
