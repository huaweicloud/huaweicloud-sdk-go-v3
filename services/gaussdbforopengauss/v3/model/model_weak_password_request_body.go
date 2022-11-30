package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WeakPasswordRequestBody struct {

	// 数据库帐号密码。
	Password string `json:"password"`
}

func (o WeakPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"WeakPasswordRequestBody", string(data)}, " ")
}
