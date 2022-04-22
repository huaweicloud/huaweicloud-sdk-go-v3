package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAppSecretReq struct {

	// 应用认证访问SECRET,未提供（字段不存在或值为null）时随机生成 - 字符集：英文字母、数字、！、@、#、$、%、+、=、点、中划线、斜线/ - 复杂度：满足大小写字母、数字、特殊字符的复杂度组合，考虑兼容性暂时可不做
	Secret *string `json:"secret,omitempty"`
}

func (o UpdateAppSecretReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppSecretReq struct{}"
	}

	return strings.Join([]string{"UpdateAppSecretReq", string(data)}, " ")
}
