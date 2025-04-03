package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PatchUserReqBody struct {

	// 概要
	Schemas []string `json:"schemas"`

	// 要执行的修改操作列表
	Operations []OperationItemDto `json:"Operations"`
}

func (o PatchUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchUserReqBody struct{}"
	}

	return strings.Join([]string{"PatchUserReqBody", string(data)}, " ")
}
