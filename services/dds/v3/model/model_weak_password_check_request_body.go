package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WeakPasswordCheckRequestBody struct {

	// 密码
	Password string `json:"password"`
}

func (o WeakPasswordCheckRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPasswordCheckRequestBody struct{}"
	}

	return strings.Join([]string{"WeakPasswordCheckRequestBody", string(data)}, " ")
}
