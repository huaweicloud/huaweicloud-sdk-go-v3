package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PatchGroupReqBody struct {

	// 概要
	Schemas []string `json:"schemas"`

	// 要执行的修改操作列表
	Operations []OperationItemDto `json:"Operations"`
}

func (o PatchGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchGroupReqBody struct{}"
	}

	return strings.Join([]string{"PatchGroupReqBody", string(data)}, " ")
}
