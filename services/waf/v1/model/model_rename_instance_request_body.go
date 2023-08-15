package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameInstanceRequestBody 独享引擎改名请求数据
type RenameInstanceRequestBody struct {

	// 独享引擎新名称
	Instancename string `json:"instancename"`
}

func (o RenameInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RenameInstanceRequestBody", string(data)}, " ")
}
