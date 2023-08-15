package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateProjectRequestBody 创建项目的请求体。
type KeystoneCreateProjectRequestBody struct {
	Project *KeystoneCreateProjectOption `json:"project"`
}

func (o KeystoneCreateProjectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateProjectRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProjectRequestBody", string(data)}, " ")
}
